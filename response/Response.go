package response

import "github.com/gin-gonic/gin"

func Json(code int, message string, obj interface{}) {
	var c *gin.Context
	c.JSON(code, gin.H{"code": code, "message": message, "data": obj})
}

func App(app string, message string, author string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"app_name": app, "version": message, "author": author})
	}
}
