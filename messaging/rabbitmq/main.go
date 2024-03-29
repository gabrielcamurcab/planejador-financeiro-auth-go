package rabbitmq

import "github.com/streadway/amqp"

type ConnectionOptions struct {
	URL      string
	Exchange string
	Queue    string
}

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
	handler func([]byte)
}

func NewConsumer(opts ConnectionOptions) (*Consumer, error) {
	conn, err := amqp.Dial(opts.URL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = channel.ExchangeDeclare(
		opts.Exchange,
		amqp.ExchangeFanout,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	queue, err := channel.QueueDeclare(
		opts.Queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		conn:    conn,
		channel: channel,
		queue:   queue.Name,
	}, nil
}

func (c *Consumer) RegisterHandler(handler func([]byte)) {
	c.handler = handler
}

func (c *Consumer) Start() error {
	msgs, err := c.channel.Consume(
		c.queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			c.handler(msg.Body)
		}
	}()

	return nil
}
