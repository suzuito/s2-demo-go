package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/s2-demo-go/sgin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/fn/point/PointFromLatLng", sgin.FnPointFromLatLng())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
