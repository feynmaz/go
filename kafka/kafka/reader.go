package kafka

import (
	"context"
	"fmt"
	"log"

	kafkago "github.com/segmentio/kafka-go"
)

type Reader struct {
	Reader *kafkago.Reader
}

func NewKafkaReader() *Reader {
	reader := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: []string{"localhost:19092"},
		Topic:   "topic_1",
		GroupID: "groupz",
	})

	return &Reader{
		Reader: reader,
	}
}

func (r *Reader) FetchMessage(ctx context.Context, messages chan<- kafkago.Message) error {
	for {
		message, err := r.Reader.FetchMessage(ctx)
		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case messages <- message:
			log.Printf("messsage fetched and sent to the channel: %s \n", message.Value)
		}
	}
}

func (r *Reader) CommitMessage(ctx context.Context, messageCommitChan <-chan kafkago.Message) error {
	for {
		select {
		case <-ctx.Done():
		case msg := <-messageCommitChan:
			err := r.Reader.CommitMessages(ctx, msg)
			if err != nil {
				return fmt.Errorf("failed to commit message: %w", err)
			}
			log.Printf("message committed: %s", msg.Value)
		}
	}
}
