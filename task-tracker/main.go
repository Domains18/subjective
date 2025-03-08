package main;

import(
	"github.com/gin-gonic/gin"
	"log"
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

	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	});

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	if err := router.Run(":8080"); err != nil{
		log.Fatal("Unable to start the server: %v", err);
	}
}