package main

import (
	"crud_go/controller"
	"crud_go/util"

	"github.com/gin-gonic/gin"
)

func main() {

	server := controller.Server{
		DB: util.InitDB(),
	}
	router := gin.Default()

	router.LoadHTMLGlob("view/*")

	router.GET("/", server.GetStorePage)
	router.GET("/new", server.NewStorePage)

	router.POST("/new", server.CreateStoreHandler)
	router.POST("/delete/:id", server.DeleteStoreHandler)

	router.Run(":8080")
}
