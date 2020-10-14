package http

import (
	err "github.com/MuhammadSuryono1997/framework-okta/base/error"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type serviceInfo struct {
	AppName string `json:"app_name"`
	Version string `json:"version"`
	Author  string `json:"author"`
}

func ServiceInfo(app string, message string, author string) gin.HandlerFunc {
	return func(c *gin.Context) {
		response := BaseResponse{
			IsSuccess: true,
			Error:     nil,
			Data: serviceInfo{
				AppName: app,
				Version: message,
				Author:  author,
			},
		}
		c.JSON(200, response)
	}
}

func CreateHttpServer() *gin.Engine {

	errorEnv := godotenv.Load()
	if errorEnv != nil {
		log.Fatal("Error loading .env file")
	}

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"

		},
		MaxAge: 12 * time.Hour,
	}))

	server.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, err.NOT_FOUND.AsInvalidResponse())
	})

	server.GET("/", ServiceInfo(
		os.Getenv("APP_NAME"),
		os.Getenv("VERSION"),
		"TEAM_BACKEND_OKTAPOS"))

	return server
}
