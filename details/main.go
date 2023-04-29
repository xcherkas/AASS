// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-db/postgres"
	// "github.com/rookie-ninja/rk-gin/v2/boot"
	camunda_client_go "github.com/citilinkru/camunda-client-go/v3"
	processor "github.com/citilinkru/camunda-client-go/v3/processor"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"time"
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
		WorkerId:                  "demo-worker-products",
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
			{TopicName: "aass-product-details"},
		},
		func(ctx *processor.Context) error {
			_id, ok := ctx.Task.Variables["itemID"].Value.(string)

			fmt.Println("id: ", _id)
			if !ok {
				return ctx.HandleFailure(processor.QueryHandleFailure{
					ErrorMessage: &errMsg,
					Retries: &retries,
					RetryTimeout: &retryTimeout,
				})
				return nil
			}

			id, _ := strconv.Atoi(_id)
			prod := &Product{}
			res := dbInstance.Where("id = ?", id).First(prod)

			if res.Error != nil {
				if errors.Is(res.Error, gorm.ErrRecordNotFound)	{
					return ctx.HandleFailure(processor.QueryHandleFailure{
						ErrorMessage: &errMsg,
						Retries: &retries,
						RetryTimeout: &retryTimeout,
					})
				}
				return ctx.HandleFailure(processor.QueryHandleFailure{
					ErrorMessage: &errMsg,
					Retries: &retries,
					RetryTimeout: &retryTimeout,
				})
			}

			// mapping results from prod array to camunda variables map
			variables := make(map[string]camunda_client_go.Variable)
			variables["productName"] = camunda_client_go.Variable{
				Value: prod.Title,
				Type: "string",
			}
			variables["productPrice"] = camunda_client_go.Variable{
				Value: prod.Price,
				Type: "double",
			}
			variables["productDesc"] = camunda_client_go.Variable{
				Value: prod.Desc,
				Type: "string",
			}
			variables["productType"] = camunda_client_go.Variable{
				Value: prod.Type,
				Type: "string",
			}
			variables["productImg"] = camunda_client_go.Variable{
				Value: prod.Img,
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

func DetailShower(ctx *gin.Context) {
	id := ctx.Request.URL.Query().Get("itemID")
	prod := &Product{}
	res := dbInstance.Where("id = ?", id).First(prod)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound)	{
			ctx.JSON(http.StatusNotFound, "No such item")
			return
		}
		ctx.JSON(http.StatusInternalServerError, res.Error)
		return
	}

	ctx.JSON(http.StatusOK, prod)
}
