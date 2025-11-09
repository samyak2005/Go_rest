package main

import (
	"example/rest/db"
	"example/rest/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	db.InitDB()
	server := gin.Default()
	routes.InitRoutes(server)
	server.Run(":8080")
}

// func getEvents(c *gin.Context){
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Hello World",
// 	})
// }

