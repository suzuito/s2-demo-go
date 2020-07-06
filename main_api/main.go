package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/s2-demo-go/sgin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("ALLOW_ORIGINS")},
		AllowMethods: []string{os.Getenv("ALLOW_METHODS")},
		AllowHeaders: []string{os.Getenv("ALLOW_HEADERS")},
	}))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/fn/point/all_expression", sgin.PointAllExpression())
	r.GET("/edge/new", sgin.Edge())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
