package sgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/geo/s2"
)

// CellUnionRegionCoverer ...
func CellUnionRegionCoverer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := struct {
			Region   []LatLngLiteral `json:"region"`
			MinLevel int             `json:"minLevel"`
			MaxLevel int             `json:"maxLevel"`
			LevelMod int             `json:"levelMod"`
			MaxCells int             `json:"maxCells"`
		}{}
		if err := ctx.BindJSON(&body); err != nil {
			abortError(ctx, err)
			return
		}
		region := NewS2LoopFromLatLngLiteralArray(body.Region)
		coverer := s2.RegionCoverer{
			MinLevel: body.MinLevel,
			MaxLevel: body.MaxLevel,
			MaxCells: body.MaxCells,
			LevelMod: body.LevelMod,
		}
		ctx.IndentedJSON(http.StatusOK, struct {
			CellUnion         *[]*CellLiteral
			Covering          *[]*CellLiteral
			FastCovering      *[]*CellLiteral
			InteriorCellUnion *[]*CellLiteral
			InteriorCovering  *[]*CellLiteral
		}{
			CellUnion:         NewCellLiteralsFromCellIDs(coverer.CellUnion(region)),
			Covering:          NewCellLiteralsFromCellIDs(coverer.Covering(region)),
			FastCovering:      NewCellLiteralsFromCellIDs(coverer.FastCovering(region)),
			InteriorCellUnion: NewCellLiteralsFromCellIDs(coverer.InteriorCellUnion(region)),
			InteriorCovering:  NewCellLiteralsFromCellIDs(coverer.InteriorCovering(region)),
		})
	}
}
