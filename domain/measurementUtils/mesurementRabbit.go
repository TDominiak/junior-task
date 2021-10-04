package measurementUtils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/influxdata/influxdb-client-go/api"
	"github.com/streadway/amqp"
)

type MeasurementRabbitExchanger struct {
	Writer  api.WriteAPIBlocking
	channel *amqp.Channel
}

func NewMeasurementRabbitExchanger() (MeasurementExchanger, error) {
	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"measurements",
		"topic",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to decalare exchanger", err)
		return nil, err
	}

	_, err = ch.QueueDeclare(
		"measurements", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
		return nil, err
	}

	err = ch.QueueBind(
		"measurements",
		"*",
		"measurements",
		false,
		nil)

	if err != nil {
		log.Fatalf("%s: %s", "Failed bind queue", err)
		return nil, err
	}

	return &MeasurementRabbitExchanger{Writer: NewClientInflux(), channel: ch}, nil
}

func (ms *MeasurementRabbitExchanger) Publish(id string, value float64) error {

	err := ms.channel.Publish(
		"measurements", // exchange
		id,             // routing key. Should be: tenant.device_id ??????
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("%f", value)),
		})

	if err != nil {
		log.Fatalf("%s: %s", "Failed to publish a message", err)
		return err
	}
	log.Printf("Measurements send from device %v", id)

	return nil
}

func (ms *MeasurementRabbitExchanger) StartConsuming() error {

	msgs, err := ms.channel.Consume(
		"measurements", // queue
		"",             // consumer
		true,           // auto ack
		false,          // exclusive
		false,          // no local
		false,          // no wait
		nil,            // args
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to register a consumer", err)
		return err
	}
	go func() {
		log.Printf("Start consuming...")
		for d := range msgs {
			value, err := strconv.ParseFloat(string(d.Body), 64)
			if err != nil {
				log.Print(err.Error())
				continue
			}
			savePoint(Measurement{Id: d.RoutingKey, Value: value}, ms.Writer)
			log.Printf("Measurements recaived from device %v", d.RoutingKey)
		}
	}()

	return nil
}
