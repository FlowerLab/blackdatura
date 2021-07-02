// +build bd_all bd_kafka kafka

package blackdatura

import (
	"net/url"
	"time"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

// DefaultKafka create a default kafka sink instance
func DefaultKafka(topic string, brokers []string) KafkaSink {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.NoResponse
	config.Producer.Retry.Max = 3
	return getKafkaSink(brokers, topic, config)
}

// Kafka create kafka sink instance
func Kafka(topic string, brokers []string, config *sarama.Config) KafkaSink {
	if config == nil || config.Validate() != nil {
		panic("config error")
	}
	return getKafkaSink(brokers, topic, config)
}

type KafkaSink struct {
	kafkaProducer sarama.SyncProducer
	topic         string
}

func getKafkaSink(brokers []string, topic string, config *sarama.Config) KafkaSink {
	producerInst, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}
	return KafkaSink{
		kafkaProducer: producerInst,
		topic:         topic,
	}
}

// Sink instance
func (p KafkaSink) Sink(*url.URL) (zap.Sink, error) {
	return p, nil
}

// Write implement zap.Sink func Write
func (p KafkaSink) Write(b []byte) (n int, err error) {
	_, _, err = p.kafkaProducer.SendMessage(&sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.StringEncoder(time.Now().String()),
		Value: sarama.ByteEncoder(b),
	})
	return len(b), err
}

// Sync implement zap.Sink func Sync
func (KafkaSink) Sync() error { return nil }

// Close implement zap.Sink func Close
func (KafkaSink) Close() error { return nil }

// String implement zap.Sink func Close
func (KafkaSink) String() string { return "kafka" }
