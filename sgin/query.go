package sgin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/geo/s2"
	"github.com/suzuito/common-go/cgin"
)

func queryLatLng(
	ctx *gin.Context,
	keyNameLat, keyNameLng string,
) (s2.LatLng, error) {
	lat := cgin.DefaultQueryAsFloat64(ctx, keyNameLat, -10000)
	lng := cgin.DefaultQueryAsFloat64(ctx, keyNameLng, -10000)
	if lat <= -10000 || lng <= -10000 {
		return s2.LatLng{}, &httpError{
			http.StatusBadRequest,
			fmt.Sprintf("%s %s must be larger than 0", keyNameLat, keyNameLng),
			nil,
		}
	}
	return s2.LatLngFromDegrees(lat, lng), nil
}
