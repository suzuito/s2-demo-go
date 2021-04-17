package main

import (
	"fmt"

	"github.com/golang/geo/s2"
)

func main() {
	var l s2.LatLng

	// 良い例
	l = s2.LatLngFromDegrees(35.6938, 139.7034) // 新宿駅の座標
	fmt.Printf("緯度=%f 経度=%f（ラジアン表現）\n", l.Lat, l.Lng)
	fmt.Printf("緯度=%f 経度=%f（度数表現）\n", l.Lat.Degrees(), l.Lng.Degrees())
	fmt.Printf("IsValid=%v\n", l.IsValid())

	// 注意
	// String関数は度数表現を出力する
	// （実際には、Lat Lng はラジアン表現の値が格納されている）
	fmt.Printf("String %s\n", l.String())
	fmt.Printf("String %s\n", l)

	// ダメな例
	l = s2.LatLng{Lat: 35.6938, Lng: 139.7034} // 新宿駅の座標を誤って設定
	fmt.Printf("緯度=%f 経度=%f（ラジアン表現）誤\n", l.Lat, l.Lng)
	fmt.Printf("緯度=%f 経度=%f（度数表現）誤\n", l.Lat.Degrees(), l.Lng.Degrees())
	fmt.Printf("IsValid=%v\n", l.IsValid())
}
