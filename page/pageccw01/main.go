package main

import (
	"fmt"

	"github.com/golang/geo/s2"
	"github.com/suzuito/s2-demo-go/s2geojson"
)

func dirToName(d s2.Direction) string {
	if d == s2.Clockwise {
		return "時計回り"
	}
	if d == s2.CounterClockwise {
		return "反時計回り"
	}
	return "不定"
}

func main() {
	a := s2.PointFromLatLng(s2.LatLngFromDegrees(43.0618, 141.3545)) // 札幌
	b := s2.PointFromLatLng(s2.LatLngFromDegrees(35.6804, 139.7690)) // 東京
	c := s2.PointFromLatLng(s2.LatLngFromDegrees(33.5902, 130.4017)) // 福岡

	fmt.Printf("東京->札幌->福岡 %s\n", dirToName(s2.RobustSign(b, a, c)))
	fmt.Printf("東京->福岡->札幌 %s\n", dirToName(s2.RobustSign(b, c, a)))

	o := s2.PointFromLatLng(s2.LatLngFromDegrees(43.1332, 131.9113)) // ウラジオストク
	fmt.Printf("ウラジオストクからみた 東京->札幌->福岡 OrderedCCW: %v\n",
		s2.OrderedCCW(b, a, c, o))
	fmt.Printf("東京からみた ウラジオストク->札幌->福岡 OrderedCCW: %v\n",
		s2.OrderedCCW(o, a, c, b))

	s2geojson.Print(&[]s2.Point{b, c, a, o})
}
