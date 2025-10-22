package broker

import (
	"context"
	"crypto/tls"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/jmontesinos91/oevents"
	ologger "github.com/jmontesinos91/ologs/logger"

	"github.com/sirupsen/logrus"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/plain"
)

// Client Represents a Kafka Client
type Client struct {
	Conn *kgo.Client
	log  *ologger.ContextLogger
}

// Connect Creates a new secure connection to Kafka and returns its reference
func Connect(config OBrokerConfig, logger *ologger.ContextLogger) (*Client, error) {

	// TLS configuration
	tlsDialer := &tls.Dialer{NetDialer: &net.Dialer{Timeout: 10 * time.Second}}

	// General Settings
	opts := []kgo.Opt{
		kgo.ClientID(config.ClientName),
		kgo.SeedBrokers(config.Servers),
		kgo.SASL(plain.Auth{
			User: config.User,
			Pass: config.Password,
		}.AsMechanism()),
		kgo.Dialer(tlsDialer.DialContext),
		kgo.RequiredAcks(kgo.AllISRAcks()),
		kgo.RecordRetries(3),
		//kgo.WithLogger(kzap.New(logger.Desugar())),
	}

	// Consumer Settings
	if config.ConsumerEnabled {
		opts = append(opts, kgo.ConsumerGroup(config.ConsumerGroupName),
			kgo.ConsumeTopics(config.ConsumeFromTopics...),
			kgo.MaxConcurrentFetches(3),
			kgo.SessionTimeout(40*time.Second),
			kgo.FetchMaxWait(15*time.Second),
			kgo.FetchMaxBytes(1048576), // 1MB
			kgo.DisableAutoCommit())
	}

	// Client configuration
	cl, err := kgo.NewClient(opts...)

	if err != nil {
		logger.WithContext(
			logrus.ErrorLevel,
			"Connect",
			"Kafka connection failed ",
			ologger.Context{},
			err)
		return nil, err
	}

	kc := &Client{
		Conn: cl,
		log:  logger,
	}

	return kc, nil
}

// ConnectInsecure Creates a new insecure connection to Kafka and returns its reference.
// Only use it for development
func ConnectInsecure(config OBrokerConfig, logger *ologger.ContextLogger) (*Client, error) {
	// General Settings
	opts := []kgo.Opt{
		kgo.ClientID(config.ClientName),
		kgo.SeedBrokers(config.Servers),
		kgo.RequiredAcks(kgo.AllISRAcks()),
		kgo.RecordRetries(3),
		//kgo.WithLogger(kzap.New(logger.Desugar())),
	}

	// Consumer Settings
	if config.ConsumerEnabled {
		opts = append(opts, kgo.ConsumerGroup(config.ConsumerGroupName),
			kgo.ConsumeTopics(config.ConsumeFromTopics...),
			kgo.MaxConcurrentFetches(3),
			kgo.SessionTimeout(30*time.Second),
			kgo.FetchMaxWait(5*time.Second),
			kgo.FetchMaxBytes(1048576), // 1MB
			kgo.DisableAutoCommit())
	}

	// Client configuration
	cl, err := kgo.NewClient(opts...)

	if err != nil {
		logger.WithContext(
			logrus.ErrorLevel,
			"Connect Insecure",
			"Kafka connection failed.",
			ologger.Context{},
			err)
		return nil, err
	}

	kc := &Client{
		Conn: cl,
		log:  logger,
	}

	return kc, nil
}

// Publish Produces a new message to the event stream
func (kc *Client) Publish(ctx context.Context, topic string, events ...oevents.OmniViewEvent) bool {

	var records []*kgo.Record

	// Convert to Kafka records
	for _, v := range events {
		data := v.ToJSON()
		record := &kgo.Record{Topic: topic, Key: []byte(v.ID), Value: []byte(data)}
		records = append(records, record)
	}

	// Publish records and wait for response
	resp := kc.Conn.ProduceSync(ctx, records...)

	// On Error return false
	if err := resp.FirstErr(); err != nil {
		kc.log.Log(
			logrus.WarnLevel,
			"Publish",
			"Failed to publish message. "+err.Error()+"")
		return false
	}

	kc.log.Log(
		logrus.InfoLevel,
		"Publish",
		"Broker stats: "+strconv.Itoa(len(records))+" events published successfully")

	return true
}

// Subscribe Consumes messages from the pre-configured topics during the connection step.
// Internally, messages are consumed within its own go-routine and made available to the caller
// through a bounded channel
func (kc *Client) Subscribe(ctx context.Context, maxRecords int, workerChannel chan<- OmniViewMessage) {
	// Consumer works on its own routine
	go func() {
		var wg sync.WaitGroup

		// Main, infinite loop for polling
		for {
			kc.log.Log(
				logrus.InfoLevel,
				"Subscribe",
				"Kafka - Polling for new messages...")

			// Will block until messages become available
			fetches := kc.Conn.PollRecords(ctx, maxRecords)

			if fetches.IsClientClosed() {
				kc.log.Log(
					logrus.WarnLevel,
					"Subscribe",
					"Kafka client closed")
				return
			}

			// Seeks for any potential errors in messages
			fetches.EachError(func(t string, p int32, err error) {
				kc.log.WithContext(
					logrus.ErrorLevel,
					"Subscribe",
					"fetch err topic "+t+" "+
						"partition "+strconv.Itoa(int(p))+": ",
					ologger.Context{}, err)
			})

			kc.log.Log(
				logrus.InfoLevel,
				"Subscribe",
				"Pulled "+strconv.Itoa(len(fetches.Records()))+" messages ready to be processed")

			// Iterates over messages
			fetches.EachRecord(func(r *kgo.Record) {
				// Parse into a OmniView Event
				te, err := oevents.ParseEvent(r.Value)
				if err != nil {
					kc.log.Log(
						logrus.WarnLevel,
						"Subscribe",
						"Error while parsing event"+err.Error()+". Payload: "+string(r.Value)+"")
				} else {
					wg.Add(1)

					msg := OmniViewMessage{
						Event: *te,
						Ack:   &wg,
					}

					// Forward through the channel
					workerChannel <- msg
				}
			})

			// Block until the batch is completed
			wg.Wait()

			// Commit all pending offsets
			if err := kc.Conn.CommitUncommittedOffsets(ctx); err != nil {
				kc.log.Log(logrus.InfoLevel, "Subscribe", "commit records failed: "+err.Error()+"")
				continue
			}
		}
	}()
}

// Close Shuts down Kafka connection
func (kc *Client) Close() {
	kc.Conn.Close()
}
