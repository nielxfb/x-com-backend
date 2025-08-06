package main

import (
	"api-gateway/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	{
		user := r.Group("/user")
		user.GET("/login", handler.Login)
	}

	r.Run()
}