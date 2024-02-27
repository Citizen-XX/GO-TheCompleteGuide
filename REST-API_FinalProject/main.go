package main

import (
	"restapimod/REST-API_FinalProject/db"
	"restapimod/REST-API_FinalProject/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") // localhost:8080
}