import (
	"context"
    "encoding/json"
    "github.com/streadway/amqp"
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {
	jsonData, err := json.Marshal(val)
	if err != nil {
		return err
	}

	msg := amqp.PublishWithContext(
		context.Background(),
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: jsonData,
		},
	)
	if err := nil {
		return err
	}

	return nil
}