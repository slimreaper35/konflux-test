package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slimreaper35/konflux-test/database"
	"github.com/slimreaper35/konflux-test/routes"
	"github.com/slimreaper35/konflux-test/utils"
)

func main() {
	database.InitDatabase()
	gin.SetMode(gin.ReleaseMode)

	var server = gin.Default()
	registerRoutes(server)

	println("Server is running on port 8080")

	var addr = "0.0.0.0:8080"
	server.Run(addr)
}

func registerRoutes(server *gin.Engine) {
	// Root
	server.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Welcome to REST API built in Konflux")
	})

	// Events
	server.GET("/events", routes.GetAllEventsHandler)
	server.GET("/events/:id", routes.GetOneEventHandler)

	var authenticated = server.Group("/")
	authenticated.Use(authenticationMiddleware)

	authenticated.POST("/events", routes.CreateEventHandler)
	authenticated.PUT("/events/:id", routes.UpdateEventHandler)
	authenticated.DELETE("/events/:id", routes.DeleteEventHandler)

	// Registrations
	authenticated.POST("/events/:id/register", routes.RegisterForEventHandler)
	authenticated.DELETE("/events/:id/register", routes.UnregisterFromEventHandler)

	// Users
	server.POST("/signup", routes.SignUpHandler)
	server.POST("/login", routes.LoginHandler)
}

func authenticationMiddleware(context *gin.Context) {
	var tokenString = context.Request.Header.Get("Authorization")
	if tokenString == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	userID, err := utils.VerifyJWT(tokenString)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	context.Set("userID", userID)
	context.Next()
}
