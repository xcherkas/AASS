// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

package main

import (
  "github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-db/postgres"
	"gorm.io/gorm"

	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
  "strings"
  "encoding/json"
  "time"

	"github.com/Shopify/sarama"
	"strconv"
)

var dbInstance *gorm.DB


func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Bootstrap
	boot.Bootstrap(context.TODO())

	pgEntry := rkpostgres.GetPostgresEntry("aass-db")
	dbInstance = pgEntry.GetDB("aass")
	if !dbInstance.DryRun {
		dbInstance.AutoMigrate(&Product{})
	}

	// Create a Kafka producer
	config := sarama.NewConfig()
	producer, err := sarama.NewAsyncProducer([]string{"kafka:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating producer: %s", err.Error())
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Error closing producer: %s", err.Error())
		}
	}()

	// Create a Kafka consumer group
	consumerGroup, err := sarama.NewConsumerGroup([]string{"kafka:9092"}, "eshop__consumers__details", config)
	if err != nil {
		log.Fatalf("Error creating consumer group: %s", err.Error())
	}
	defer func() {
		if err := consumerGroup.Close(); err != nil {
			log.Fatalf("Error closing consumer group: %s", err.Error())
		}
	}()

	// Create a channel to receive OS signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Start the consumer group
	go func() {
		for {
			err := consumerGroup.Consume(context.Background(), []string{"eshop"}, &MyConsumer{producer: producer})
			if err != nil {
				log.Fatalf("Error consuming messages: %s", err.Error())
			}
		}
	}()

	boot.WaitForShutdownSig(context.TODO())
}

// MyConsumer implements the sarama.ConsumerGroupHandler interface
type MyConsumer struct {
	producer sarama.AsyncProducer
}

func (c *MyConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *MyConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func Buildup(name string, id string, value string) string {
  return name + ":@?@:" + id + ":@?@:" + value
}

func (c *MyConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
    msg := string(message.Value)

    if !strings.HasPrefix(msg, "ms-details") {
      fmt.Println("Skipping...\n")
      session.MarkMessage(message, "")
      continue
    }

    preparsed := strings.Split(msg, ":@?@:")

    var result Input
    pload := []byte(preparsed[2])
    err := json.Unmarshal(pload, &result)

    var prod Product
		res := dbInstance.Where("id = ?", strconv.Itoa(result.Id)).First(&prod)


    if res.Error != nil {
      c.producer.Input() <- &sarama.ProducerMessage{Topic: "eshop", Value: sarama.StringEncoder(Buildup("viewer",preparsed[1],"{\"error\": \"Error occured\"}"))}
      session.MarkMessage(message, "")
      continue
    }

    jsoned, err := json.Marshal(prod)

    if err != nil {
      c.producer.Input() <- &sarama.ProducerMessage{Topic: "eshop", Value: sarama.StringEncoder(Buildup("viewer",preparsed[1],"{\"error\": \"Error occured\"}"))}
      session.MarkMessage(message, "")
      continue
    }

		c.producer.Input() <- &sarama.ProducerMessage{Topic: "eshop", Value: sarama.StringEncoder(Buildup("viewer",preparsed[1], string(jsoned)))}
    session.MarkMessage(message, "")
	}
	return nil
}

type Input struct {
  Id int `json:"id"`
}

type Base struct {
  CreatedAt time.Time      `yaml:"-" json:"-"`
  UpdatedAt time.Time      `yaml:"-" json:"-"`
  DeletedAt gorm.DeletedAt `yaml:"-" json:"-" gorm:"index"`
}

type Product struct {
  Base
  Id   int    `yaml:"id" json:"id" gorm:"primaryKey"`
  Title string `yaml:"title" json:"title"`
  Desc string `yaml:"desc" json:"desc"`
  Price float32 `yaml:"price" json:"price"`
  Img string `yaml:"img" json:"img"`
  Type string `yaml:"type" json:"type"`
}
