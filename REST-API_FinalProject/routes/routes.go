package routes

import (
	"restapimod/REST-API_FinalProject/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events",getEvents) //GET, POST, PUT, DELETE, PATCH
	server.GET("/events/:id",getEvent)	// /events/1, /events/5

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events",createEvent)
	authenticated.PUT("/events/:id",updateEvent)
	authenticated.DELETE("/events/:id",deleteEvent)
	authenticated.POST("/events/:id/register",registerForEvent)
	authenticated.DELETE("/events/:id/register",cancelRegistration)

	server.POST("/signup",signup)
	server.POST("/login",login)
}