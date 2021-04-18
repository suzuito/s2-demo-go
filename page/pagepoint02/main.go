package main

import (
	"fmt"

	"github.com/golang/geo/s2"
)

func main() {
	l1 := s2.LatLngFromDegrees(35.6938, 139.7034) // 新宿駅の座標

	p1 := s2.PointFromLatLng(l1)
	fmt.Printf("緯度=%f 経度=%f（ラジアン表現）\n", l1.Lat, l1.Lng)
	fmt.Printf("緯度=%f 経度=%f（度数表現）\n", l1.Lat.Degrees(), l1.Lng.Degrees())
	fmt.Printf("Point x=%f y=%f z=%f\n", p1.X, p1.Y, p1.Z)
	fmt.Printf("Point 3 次元ベクトルの長さ=%f\n", p1.Norm())

	p2 := s2.PointFromCoords(10000, 0, 0) // 入力されるベクトルは長さが 10000
	fmt.Printf("Point x=%f y=%f z=%f\n", p2.X, p2.Y, p2.Z)
	fmt.Printf("Point 3 次元ベクトルの長さ=%f\n", p2.Norm())
}
