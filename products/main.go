// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-db/postgres"
	"github.com/rookie-ninja/rk-gin/v2/boot"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var dbInstance *gorm.DB

func main() {

	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, dbport, user, password, dbname)
	//
	//db, err := sql.Open("postgres", psqlInfo)
	//
	//if err != nil {
	//	panic(err)
	//}
	//err = db.Ping()
	//
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Established a successful connection!")

	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Register handler
	entry := rkgin.GetGinEntry("rk-demo")
	entry.Router.GET("/v1/products", ProductList)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	pgEntry := rkpostgres.GetPostgresEntry("aass-db")
	dbInstance = pgEntry.GetDB("aass")
	if !dbInstance.DryRun {
		dbInstance.AutoMigrate(&Product{})
	}

	boot.WaitForShutdownSig(context.TODO())
}

// FilterList Gets all products and return them as a list
func ProductList(ctx *gin.Context) {
	var prod []Product
	res := dbInstance.Find(&prod)

	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, res.Error)
		return
	}

	ctx.JSON(http.StatusOK, prod)
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
