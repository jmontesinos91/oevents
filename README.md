# OmniView + Events = OEvents

A Go module providing the foundations for Omniview events

## Install

```
$ go get -u git.omnicloud.mx/omnicloud/development/go-modules/oevents
```

## Run locally
### First step
First you need to run docker-compose.yaml
### Second step
After already run, if it's the first time that you are runing this project
you need to create the topic.
#### Topic Creation
You need to open a terminal in you kafka service and run the next command
```
/opt/bitnami/kafka/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --topic {topic_name} --partitions 3 --replication-factor 1
```
To validate that already created and have all the configuration that you expected you need to use  the next command

#### Get all the Topics
```
/opt/bitnami/kafka/bin/kafka-topics.sh --list --bootstrap-server localhost:9092
```

#### Describe the topic
```
/opt/bitnami/kafka/bin/kafka-topics.sh --describe --bootstrap-server localhost:9092 --topic {topic_name}
```

#### Delete Topics
```
/opt/bitnami/kafka/bin/kafka-topics.sh --delete --bootstrap-server localhost:9092 --topic {topic_name}
```

#### Subscribe to Topic and Consume message
```
/opt/bitnami/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic {topic_name}
```

## Description
This project it's a library to consume and produce message.

## Authors and acknowledgment
Jonatan Abelardo Montesinos Macias