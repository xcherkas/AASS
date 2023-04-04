// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-gin/v2/boot"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample rk-demo server.
// @termsOfService http://swagger.io/terms/

// @securityDefinitions.basic BasicAuth

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Register handler
	entry := rkgin.GetGinEntry("rk-demo")
	entry.Router.GET("/v1/details", DetailShower)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	boot.WaitForShutdownSig(context.TODO())
}

func DetailShower(ctx *gin.Context) {
	//id := ctx.Request.URL.Query().Get("itemID")
	//sqlQuery := db.Query(`SELECT itemDesc FROM aass.items WHERE itemID=$1 VALUES($1)`,id)
	ctx.JSON(http.StatusOK, &ItemDetail{
		Item: fmt.Sprintf("Displaying item description of %s", ctx.Query("itemID")),
	})
}

type ItemDetail struct {
	Item string
}
