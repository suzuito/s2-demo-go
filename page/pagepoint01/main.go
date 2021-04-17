package main

import (
	"fmt"

	"github.com/golang/geo/s2"
)

func main() {
	var latlng s2.LatLng

	// 良い例
	latlng = s2.LatLngFromDegrees(35.6938, 139.7034) // 新宿駅の座標
	fmt.Printf(
		"緯度=%f 経度=%f（ラジアン表現）\n",
		latlng.Lat,
		latlng.Lng,
	)
	fmt.Printf(
		"緯度=%f 経度=%f（度数表現）\n",
		latlng.Lat.Degrees(),
		latlng.Lng.Degrees(),
	)
	fmt.Printf(
		"IsValid=%v\n",
		latlng.IsValid(),
	)

	// ダメな例
	latlng = s2.LatLng{Lat: 35.6938, Lng: 139.7034} // 新宿駅の座標を誤って設定
	fmt.Printf(
		"緯度=%f 経度=%f（ラジアン表現）誤\n",
		latlng.Lat,
		latlng.Lng,
	)
	fmt.Printf(
		"緯度=%f 経度=%f（度数表現）誤\n",
		latlng.Lat.Degrees(),
		latlng.Lng.Degrees(),
	)
	fmt.Printf(
		"IsValid=%v\n",
		latlng.IsValid(),
	)
}
