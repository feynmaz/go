package kafka

import (
	"context"

	kafkago "github.com/segmentio/kafka-go"
)

type Writer struct {
	Writer *kafkago.Writer
}

func NewKafkaWriter() *Writer {
	writer := &kafkago.Writer{
		Addr:  kafkago.TCP("localhost:19092"),
		Topic: "topic_1",
	}
	return &Writer{
		Writer: writer,
	}
}

func (w *Writer) WriteMessages(ctx context.Context, messages chan kafkago.Message, messageCommitChan chan kafkago.Message) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-messages:
			err := w.Writer.WriteMessages(ctx, kafkago.Message{
				Value: msg.Value,
			})
			if err != nil {
				return err
			}

			select {
			case <- ctx.Done():
			case messageCommitChan <- msg:
			}
		}
	}
}
