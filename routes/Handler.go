package routes

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MuhammadSuryono1997/framework-okta/response"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Handler() *gin.Engine {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	server.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "NOT_METHODE_ALLOWED"})
	})

	server.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/version"
		server.HandleContext(c)
	})

	server.GET("/version", response.App(os.Getenv("APP_NAME"), os.Getenv("VERSION"), "TEAM_BACKEND_OKTAPOS"))

	return server
}
