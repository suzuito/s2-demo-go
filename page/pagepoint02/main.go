package main

import (
	"fmt"

	"github.com/golang/geo/s2"
)

func main() {
	l1 := s2.LatLngFromDegrees(35.6938, 139.7034) // 新宿駅の座標

	p := s2.PointFromLatLng(l1)
	fmt.Printf("緯度=%f 経度=%f（ラジアン表現）\n", l1.Lat, l1.Lng)
	fmt.Printf("緯度=%f 経度=%f（度数表現）\n", l1.Lat.Degrees(), l1.Lng.Degrees())
	fmt.Printf("Point x=%f y=%f z=%f\n", p.X, p.Y, p.Z)
	fmt.Printf("Point Norm=%f\n", p.Norm())
}
