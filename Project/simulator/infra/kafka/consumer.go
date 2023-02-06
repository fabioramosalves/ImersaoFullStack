package kafka

import (
	"os"
	"fmt"
)

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaConsumer struct{
	MsgChan chan *ckafka.Message
}

func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChan,
	}
}

func(k *KafkaConsumer) Consume(){
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers" : os.Getenv("KafkaBootStrapServers"),
		"group:id" : os.Getenv("KafkaConsumerGroupId")
	}
	c, erro := ckafka.NewConsumer(configMap)
	if err != nil{
		log.Fatal("error consuming kafka message:" + err.Error())
	}
	topics := []string{os.Getenv("KafkaReadTopic")}

	c.SubscribeTopics(topics, nil)
	fmt.Println("Kafka consumer has been starded")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MsgChan <- msg
		}
	}
}