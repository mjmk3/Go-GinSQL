package main

import (
	"GoConn/controllers"
	"GoConn/db_client"
	"github.com/gin-gonic/gin"
)

func main() {
	db_client.InitializeDBConnect()

	r := gin.Default()
	r.POST("/create", controllers.CreatePost)
	r.GET("/all", controllers.GetPosts)
	// Running server on specific port
	if err := r.Run(":5000"); err != nil {
		panic(err.Error())
	}
}
