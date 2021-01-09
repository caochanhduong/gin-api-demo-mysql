package main

import (
	"demo-gin-api-with-gomod/controllers"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

// User is a struct containing user info
func setupRouter() *gin.Engine {
	c := gin.Default()
	client := c.Group("/apx")
	{
		client.GET("/user", controllers.GetAllUser)
		client.GET("/user/:id", controllers.GetUserByID)
		client.DELETE("/user/:id", controllers.DeleteUserByID)
		client.POST("/user", controllers.CreateUser)
		client.PUT("/user/:id", controllers.UpdateUserByID)
	}
	return c
}

func main() {
	c := setupRouter()
	c.Run(":8080")
}
