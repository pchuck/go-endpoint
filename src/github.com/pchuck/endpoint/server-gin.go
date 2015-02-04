package main

import "github.com/gin-gonic/gin"

func notYetImplemented(c *gin.Context) {
	c.String(501, "not yet implemented")
}

func writeEndpoint(c *gin.Context) {
	c.String(200, "ok")
}

func readEndpoint(c *gin.Context) {
	c.String(200, "pong")
}

func main() {
	r := gin.Default()

	// Simple group: v0
	v0 := r.Group("/v0")
	{
		v0.POST("/write", writeEndpoint)
		v0.GET("/read", readEndpoint)
	}

	// Simple group: v1
	v1 := r.Group("/v1")
	{
		v1.GET("/read", notYetImplemented)
	}

	// Listen and server on 0.0.0.0:8080
	r.Run(":8081")
}
