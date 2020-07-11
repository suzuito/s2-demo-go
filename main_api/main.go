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
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-type"},
	}))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/fn/point/all_expression", sgin.PointAllExpression())
	r.GET("/edge/new", sgin.Edge())
	r.POST("/cell/new", sgin.Cell())
	r.POST("/cell/all_parents", sgin.AllParentCells())
	r.POST("/cell/children", sgin.ChildCells())
	r.POST("/cell/from_token", sgin.CellFromToken())
	r.POST("/cell_union/region_coverer", sgin.CellUnionRegionCoverer())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
