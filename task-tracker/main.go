package main

import (
	"log"

	"github.com/Domains18/subjective.git/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Task Tracker Microservice
// @version 1.0
// @description This is one of the larger util system for task tracking written in go
// @termsOfService // @termsOfService URL_ADDRESS.io/terms/

// @contact.name API Suppor// @termsOfService URL_ADDRESS.io/terms/

// @co// @termsOfService URL_ADDRESS.io/terms/

// @contact.name API Support
// @contact.url URL_ADDRESS.swagger.io/support


func main(){


	router:= gin.Default();

	config.Connect_DB();

	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	});

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	if err := router.Run(":8044"); err != nil{
		log.Printf("Unable to start the server: %v", err);
	}
	log.Printf("server succesfully started on port 8080")
}