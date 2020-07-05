package sgin

import "github.com/gin-gonic/gin"

func FnPointFromLatLng() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ll, err := queryLatLng(ctx, "lat", "lng")
		if err != nil {
			abortError(ctx, err)
		}

	}
}
