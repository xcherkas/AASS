// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-db/postgres"
	// "github.com/rookie-ninja/rk-gin/v2/boot"
	camunda_client_go "github.com/citilinkru/camunda-client-go/v3"
	processor "github.com/citilinkru/camunda-client-go/v3/processor"
	"gorm.io/gorm"
	"encoding/json"
	"strconv"
	"time"
	"fmt"
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

	client := camunda_client_go.NewClient(camunda_client_go.ClientOptions{
		EndpointUrl: "http://localhost:8080/engine-rest",
		ApiUser:     "demo",
		ApiPassword: "demo",
		Timeout:     time.Second * 10,
	})

	logger := func(err error) {
		fmt.Println(err.Error())
	}
	proc := processor.NewProcessor(client, &processor.Options{
		WorkerId:                  "demo-worker-searcher",
		LockDuration:              time.Second * 5,
		MaxTasks:                  10,
		MaxParallelTaskPerHandler: 100,
		LongPollingTimeout:        5 * time.Second,
	}, logger)

	errMsg := "Internal Server Error"
	retries := 0
	retryTimeout := 10

	proc.AddHandler(
		[]*camunda_client_go.QueryFetchAndLockTopic{
			{TopicName: "aass-search"},
		},
		func(ctx *processor.Context) error {

			var prod []Product
			queryValue, ok := ctx.Task.Variables["query"].Value.(string)

			if !ok {
				return ctx.HandleFailure(processor.QueryHandleFailure{
					ErrorMessage: &errMsg,
					Retries: &retries,
					RetryTimeout: &retryTimeout,
				})
			}

			if (queryValue == "") {
				dbInstance.Find(&prod)
			} else {
				dbInstance.Where("title LIKE ?", "%"+ queryValue +"%").Find(&prod)
			}

			// mapping results from prod array to camunda variables map
			variables := make(map[string]camunda_client_go.Variable)

			// create array of CamFormSelect
			arrProds := []CamFormSelect{}

			i := 1
			for _, p := range prod {
				fmt.Println("id => ", strconv.Itoa(p.Id), p.Id)

				arrProds = append(arrProds, CamFormSelect{
					Label: p.Title,
					Value: strconv.Itoa(p.Id),
				})

				variables["productName" + strconv.Itoa(i)] = camunda_client_go.Variable{
					Value: "# " + p.Title,
					Type: "string",
				}
				variables["productImg" + strconv.Itoa(i)] = camunda_client_go.Variable{
					Value: p.Img,
					Type: "string",
				}
				variables["productPrice" + strconv.Itoa(i)] = camunda_client_go.Variable{
					Value: p.Price,
					Type: "double",
				}
				variables["productType" + strconv.Itoa(i)] = camunda_client_go.Variable{
					Value: p.Type,
					Type: "string",
				}
				i = i + 1
			}

			jsoned, _ := json.Marshal(arrProds)

			variables["itemIds"] = camunda_client_go.Variable{
				Value: string(jsoned),
				Type: "json",
			}
			variables["query"] = camunda_client_go.Variable{
				Value: queryValue,
				Type: "string",
			}

			err := ctx.Complete(processor.QueryComplete{
				Variables: &variables,
			})

			if err != nil {
				fmt.Printf("Error set complete task %s: %s\n", ctx.Task.Id, err)
			}
			fmt.Printf("Task %s completed with\n", ctx.Task.Id)
			fmt.Println(variables)
			return nil
		},
	)

	boot.WaitForShutdownSig(context.TODO())
}

type CamFormSelect struct {
	Label string `json:"label"`
	Value string    `json:"value"`
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
