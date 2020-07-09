package sgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/geo/s2"
)

// Cell ...
func Cell() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		}{}
		if err := ctx.BindJSON(&body); err != nil {
			abortError(ctx, err)
			return
		}
		latlng := s2.LatLngFromDegrees(body.Lat, body.Lng)
		cell := s2.CellFromLatLng(latlng)
		cellModel := NewCellLiteralFromCell(cell)
		ctx.IndentedJSON(http.StatusOK, cellModel)
	}
}

// AllParentCells ...
func AllParentCells() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := struct {
			Token string `json:"id"`
		}{}
		if err := ctx.BindJSON(&body); err != nil {
			abortError(ctx, err)
			return
		}
		cellID := s2.CellIDFromToken(body.Token)
		currentLevel := cellID.Level()
		parents := []s2.Cell{}
		for {
			parentCellID := cellID.Parent(currentLevel - 1)
			parents = append(parents, s2.CellFromCellID(parentCellID))
			if parentCellID.Level() == 1 {
				break
			}
			currentLevel = parentCellID.Level()
		}
		res := []*CellLiteral{}
		for _, parent := range parents {
			res = append(res, NewCellLiteralFromCell(parent))
		}
		ctx.IndentedJSON(http.StatusOK, res)
	}
}

// ChildCells ...
func ChildCells() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := struct {
			Token string `json:"id"`
		}{}
		if err := ctx.BindJSON(&body); err != nil {
			abortError(ctx, err)
			return
		}
		res := []*CellLiteral{}
		cellID := s2.CellIDFromToken(body.Token)
		for childCellID := cellID.ChildBegin(); childCellID != cellID.ChildEnd(); childCellID = childCellID.Next() {
			res = append(res, NewCellLiteralFromCell(s2.CellFromCellID(childCellID)))
		}
		ctx.IndentedJSON(http.StatusOK, res)
	}
}

// CellFromToken ...
func CellFromToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := struct {
			Token string `json:"id"`
		}{}
		if err := ctx.BindJSON(&body); err != nil {
			abortError(ctx, err)
			return
		}
		res := NewCellLiteralFromCell(s2.CellFromCellID(s2.CellIDFromToken(body.Token)))
		ctx.IndentedJSON(http.StatusOK, res)
	}
}
