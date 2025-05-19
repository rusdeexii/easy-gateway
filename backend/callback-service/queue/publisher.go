package queue

import (
    "encoding/json"
    "fmt"
    "os"

    "github.com/streadway/amqp"
    "callback-service/model"
)

func PublishTransaction(tx model.Transaction) error {
    mqUser := os.Getenv("RABBITMQ_USER")
    mqPass := os.Getenv("RABBITMQ_PASS")
    mqHost := os.Getenv("RABBITMQ_HOST")
    mqPort := os.Getenv("RABBITMQ_PORT")

    uri := fmt.Sprintf("amqp://%s:%s@%s:%s/", mqUser, mqPass, mqHost, mqPort)

    conn, err := amqp.Dial(uri)
    if err != nil {
        return err
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        return err
    }
    defer ch.Close()

    q, err := ch.QueueDeclare("transaction.created", true, false, false, false, nil)
    if err != nil {
        return err
    }

    body, err := json.Marshal(tx)
    if err != nil {
        return err
    }

    return ch.Publish("", q.Name, false, false, amqp.Publishing{
        ContentType: "application/json",
        Body:        body,
    })
}
