package service

import (
	"encoding/json"
	"fmt"
	"log"
	"notification-service/pkg"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)


type Event struct {
	ID         string `json:"id"`
	Body       string `json:"body"`
	RetryCount int    `json:"retry_count"`
}

var processed = make(map[string]bool)
var mu sync.Mutex

func Consumer(user, password, host, port string) {

	
	url := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		user,
		password,
		host,
		port,
	)

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare("orders_queue", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	retryArgs := amqp.Table{
		"x-message-ttl":             int32(5000),
		"x-dead-letter-exchange":    "",
		"x-dead-letter-routing-key": "orders_queue",
	}

	_, err = ch.QueueDeclare("orders_retry_queue", true, false, false, false, retryArgs)
	if err != nil {
		log.Fatal(err)
	}

	_, err = ch.QueueDeclare("orders_dlq", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume("orders_queue", "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("🚀 consumer started")

	go func() {
		for d := range msgs {

			var event Event
			_ = json.Unmarshal(d.Body, &event)

			// 🧠 IDEMPOTENCY CHECK
			mu.Lock()
			if processed[event.ID] {
				log.Println("⚠️ duplicate message ignored:", event.ID)
				mu.Unlock()
				d.Ack(false)
				continue
			}
			processed[event.ID] = true
			mu.Unlock()

			fmt.Println("📩 processing:", event.Body)

			err := pkg.Process(event.Body)

			if err != nil {

				if event.RetryCount >= 3 {

					log.Println("💀 DLQ")

					ch.Publish(
						"",
						"orders_dlq",
						false,
						false,
						amqp.Publishing{
							ContentType: "application/json",
							Body:        d.Body,
						},
					)

					d.Ack(false)
					continue
				}

				event.RetryCount++

				newBody, _ := json.Marshal(event)

				ch.Publish(
					"",
					"orders_retry_queue",
					false,
					false,
					amqp.Publishing{
						ContentType: "application/json",
						Body:        newBody,
					},
				)

				d.Ack(false)
				continue
			}

			d.Ack(false)
		}
	}()

	select {}
}

