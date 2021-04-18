package main

import (
	"os"

	"github.com/golang/geo/s2"
	"github.com/suzuito/s2-demo-go/s2geojson"
)

func main() {
	sapporo := s2.PointFromLatLng(s2.LatLngFromDegrees(43.0618, 141.3545))
	tokyo := s2.PointFromLatLng(s2.LatLngFromDegrees(35.6804, 139.7690))
	fukuoka := s2.PointFromLatLng(s2.LatLngFromDegrees(33.5902, 130.4017))

	s2geojson.Fprint(
		os.Stdout,
		&[]s2.Point{
			sapporo,
			tokyo,
			fukuoka,
		},
	)
}
