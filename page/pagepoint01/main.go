package main

import (
	"fmt"

	"github.com/golang/geo/s2"
)

func main() {
	latlng := s2.LatLngFromDegrees(35.6938, 139.7034)
	fmt.Printf("緯度=%f 経度=%f （ラジアン表現）\n", latlng.Lat, latlng.Lng)
	fmt.Printf("緯度=%f 経度=%f （度数表現）\n", latlng.Lat.Degrees(), latlng.Lng.Degrees())
}
