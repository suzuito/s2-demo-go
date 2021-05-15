package main

import (
	"fmt"

	"github.com/golang/geo/r3"
	"github.com/golang/geo/s2"
)

func main() {
	// Point 構造体の生成
	// (10, 0, 0) のベクトルは長さが 1 ではない。
	// PointFromCoords ファクトリメソッドは、長さが 1 ではない
	// ベクトルが入力として与えられた場合、長さが 1 のベクトルに
	// 自動的に変換し、Point 構造体を生成する。
	p1 := s2.PointFromCoords(10, 0, 0)
	fmt.Printf("p1 (%f %f %f) 長さが 1 ですか？ %v\n", p1.X, p1.Y, p1.Z, p1.IsUnit())
	p2 := s2.PointFromCoords(10, 10, 10)
	fmt.Printf("p2 (%f %f %f) 長さが 1 ですか？ %v\n", p2.X, p2.Y, p2.Z, p2.IsUnit())

	// ダメな例 長さが 1 ではないベクトルを生成できてしまう
	p3 := s2.Point{Vector: r3.Vector{X: 10, Y: 0, Z: 0}}
	fmt.Printf("p3 (%f %f %f) 長さが 1 ですか？ %v\n", p3.X, p3.Y, p3.Z, p3.IsUnit())

	// 緯度経度から Point 構造体の生成
	p4 := s2.PointFromLatLng(s2.LatLngFromDegrees(90, 0)) // 北極
	fmt.Printf("p4 (%f %f %f) 長さが 1 ですか？ %v\n", p4.X, p4.Y, p4.Z, p4.IsUnit())
	p5 := s2.PointFromLatLng(s2.LatLngFromDegrees(-90, 0)) // 南極
	fmt.Printf("p5 (%f %f %f) 長さが 1 ですか？ %v\n", p5.X, p5.Y, p5.Z, p5.IsUnit())
}
