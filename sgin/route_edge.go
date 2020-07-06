package sgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/geo/s1"
	"github.com/golang/geo/s2"
)

// Edge ...
func Edge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ll1, err := queryLatLng(ctx, "lat1", "lng1")
		if err != nil {
			abortError(ctx, err)
			return
		}
		ll2, err := queryLatLng(ctx, "lat2", "lng2")
		if err != nil {
			abortError(ctx, err)
			return
		}
		p1 := s2.PointFromLatLng(ll1)
		p2 := s2.PointFromLatLng(ll2)
		e := s2.Edge{
			V0: p1,
			V1: p2,
		}
		ctx.IndentedJSON(
			http.StatusOK,
			struct {
				Edge              s2.Edge  `json:"edge"`
				DistanceAsAngle   s1.Angle `json:"distanceAsAngle"`
				DistanceAsDegrees float64  `json:"distanceAsDegrees"`
			}{
				Edge:              e,
				DistanceAsAngle:   p1.Distance(p2),
				DistanceAsDegrees: p1.Distance(p2).Degrees(),
			},
		)
	}
}
