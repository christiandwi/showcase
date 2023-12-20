package event

import (
	"context"
	"log"

	"github.com/christiandwi/showcase/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMq interface {
	RabbitMqQueueOpen(queueName string) amqp.Queue
	RabbitMqPublish(ctx context.Context, q amqp.Queue, body []byte) error
	RabbitMqConsume(q amqp.Queue, f func(amqp.Delivery))
}

type rabbitMq struct {
	channel *amqp.Channel
}

func RabbitMqInit(config *config.Config) rabbitMq {
	conn, err := amqp.Dial(config.RabbitMq.Url)
	if err != nil {
		log.Panic("error on dial ", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Panic("error on opening channel ", err)
	}

	return rabbitMq{
		channel: ch,
	}
}

func (r rabbitMq) RabbitMqQueueOpen(queueName string) amqp.Queue {
	q, err := r.channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Panic("error on opening queue ", err)
	}

	return q
}

func (r rabbitMq) RabbitMqPublish(ctx context.Context, q amqp.Queue, body []byte) error {
	err := r.channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		log.Panic("error on publishing message")
		return err
	}
	log.Printf(" [x] Sent %s\n", body)

	return nil
}

func (r rabbitMq) RabbitMqConsume(q amqp.Queue, f func(amqp.Delivery)) {
	msgs, err := r.channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Panic("error on consuming message")
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			f(d)
		}
	}()
	<-forever

}
