package sgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/geo/s2"
)

// PointAllExpression ...
func PointAllExpression() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ll, err := queryLatLng(ctx, "lat", "lng")
		if err != nil {
			abortError(ctx, err)
			return
		}
		ctx.IndentedJSON(
			http.StatusOK,
			struct {
				Point  s2.Point  `json:"point"`
				LatLng s2.LatLng `json:"latlng"`
			}{
				Point:  s2.PointFromLatLng(ll),
				LatLng: ll,
			},
		)
	}
}
