package main

import (
	"fmt"

	"github.com/golang/geo/s2"
	"github.com/suzuito/s2-demo-go/s2geojson"
)

func main() {
	a := s2.PointFromLatLng(s2.LatLngFromDegrees(43.0618, 141.3545)) // 札幌
	b := s2.PointFromLatLng(s2.LatLngFromDegrees(35.6804, 139.7690)) // 東京
	c := s2.PointFromLatLng(s2.LatLngFromDegrees(33.5902, 130.4017)) // 福岡

	fmt.Printf("東京->札幌->福岡 %d\n", s2.RobustSign(b, a, c))
	fmt.Printf("東京->福岡->札幌 %d\n", s2.RobustSign(b, c, a))

	// ウラジオストク
	o := s2.PointFromLatLng(s2.LatLngFromDegrees(43.1332, 131.9113))
	fmt.Printf(
		"東京->札幌->福岡 は、ウラジオストクから見たとき、反時計回りですか？ %v\n",
		s2.OrderedCCW(b, a, c, o))
	fmt.Printf(
		"ウラジオストク->札幌->福岡 は、東京から見たとき、反時計回りですか？ %v\n",
		s2.OrderedCCW(o, a, c, b))

	s2geojson.Print(&[]s2.Point{b, c, a, o}, &s2geojson.PrintGeoJSONOption{
		StyleHeight: "500px",
		Zoom:        5,
		Center: s2geojson.PrintGeoJSONOptionLatLng{
			Lat: 40.2048,
			Lng: 138.2529,
		},
	})
}
