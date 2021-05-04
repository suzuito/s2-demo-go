package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/s2-demo-go/inject"
	"github.com/suzuito/s2-demo-go/sgin"
)

func main() {
	r := gin.Default()
	env, genGCP, err := inject.NewImplement()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
	sgin.NewRoute(r, env, genGCP)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// func main() {
// 	r := gin.Default()
// 	r.Use(cors.New(cors.Config{
// 		AllowOrigins: []string{os.Getenv("ALLOWED_ORIGINS")},
// 		AllowMethods: []string{"GET", "POST"},
// 		AllowHeaders: []string{"Content-type"},
// 	}))
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	r.GET("/fn/point/all_expression", sgin.PointAllExpression())
// 	r.GET("/edge/new", sgin.Edge())
// 	r.POST("/cell/new", sgin.Cell())
// 	r.POST("/cell/all_parents", sgin.AllParentCells())
// 	r.POST("/cell/children", sgin.ChildCells())
// 	r.POST("/cell/from_token", sgin.CellFromToken())
// 	r.POST("/cell_union/region_coverer", sgin.CellUnionRegionCoverer())
// 	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }
