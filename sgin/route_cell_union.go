package sgin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/geo/s2"
	geojson "github.com/paulmach/go.geojson"
	"github.com/suzuito/s2-demo-go/data"
	"github.com/suzuito/s2-demo-go/s2geojson"
)

// CellUnionRegionCoverer ...
func CellUnionRegionCoverer() gin.HandlerFunc {
	shinjuku, _ := geojson.UnmarshalGeometry([]byte(data.Shinjuku))
	region := s2geojson.NewS2LoopFromGeometry(shinjuku.Polygon[0])
	return func(ctx *gin.Context) {
		body := struct {
			MaxLevel int `json:"maxLevel"`
			LevelMod int `json:"levelMod"`
			MaxCells int `json:"maxCells"`
		}{}
		if err := ctx.BindJSON(&body); err != nil {
			abortError(ctx, err)
			return
		}
		coverer := s2.RegionCoverer{
			MinLevel: 10,
			MaxLevel: body.MaxLevel,
			MaxCells: body.MaxCells,
			LevelMod: body.LevelMod,
		}
		fmt.Println(coverer)
		ctx.IndentedJSON(http.StatusOK, struct {
			Region            *[]LatLngLiteral
			CellUnion         *[]*CellLiteral
			Covering          *[]*CellLiteral
			FastCovering      *[]*CellLiteral
			InteriorCellUnion *[]*CellLiteral
			InteriorCovering  *[]*CellLiteral
		}{
			Region:            NewLatLngLiteralArrayFromS2Loop(region),
			CellUnion:         NewCellLiteralsFromCellIDs(coverer.CellUnion(region)),
			Covering:          NewCellLiteralsFromCellIDs(coverer.Covering(region)),
			FastCovering:      NewCellLiteralsFromCellIDs(coverer.FastCovering(region)),
			InteriorCellUnion: NewCellLiteralsFromCellIDs(coverer.InteriorCellUnion(region)),
			InteriorCovering:  NewCellLiteralsFromCellIDs(coverer.InteriorCovering(region)),
		})
	}
}
