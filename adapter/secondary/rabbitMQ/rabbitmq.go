package rabbitmq

import (
	"context"
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/witalok2/test-dev-golang-api/config"
)

type QueueClient interface {
	PublishMessage(ctx context.Context, message []byte) error
}

type queueClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue

	queueName, routingKey, exchangeName string
}

func NewQueueClient(env config.Queue) (QueueClient, error) {
	conn, err := amqp.Dial(env.URI)
	if err != nil {
		return nil, errors.New("failed to connect to RabbitMQ")
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, errors.New("failed to open a channel")
	}

	err = channel.ExchangeDeclare(
		env.ExchangeName, // name
		env.ExchangeType, // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		return nil, errors.New("failed to declare a queue")
	}

	queue, err := channel.QueueDeclare(
		env.QueueName, // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return nil, errors.New("failed to declare a queue")
	}

	err = channel.QueueBind(env.QueueName, env.RoutingKey, env.ExchangeName, false, nil)
	if err != nil {
		return nil, errors.New("failed to declare exchange")
	}

	return &queueClient{
		conn:    conn,
		channel: channel,
		queue:   queue,

		queueName:    env.QueueName,
		routingKey:   env.RoutingKey,
		exchangeName: env.ExchangeName,
	}, nil
}

func (q *queueClient) PublishMessage(ctx context.Context, message []byte) error {
	err := q.channel.PublishWithContext(
		ctx,
		q.exchangeName, // exchange
		q.routingKey,   // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         message,
		},
	)
	if err != nil {
		return errors.New("failed to publish a message")
	}

	return nil
}
