package main

import (
	"fmt"

	"github.com/golang/geo/s2"
)

func main() {
	l1 := s2.LatLngFromDegrees(35.6896, 139.7006) // 新宿駅
	l2 := s2.LatLngFromDegrees(34.7025, 135.4960) // 大阪駅
	l3 := s2.LatLngFromDegrees(35.4660, 139.6221) // 横浜駅

	// 新宿駅と横浜駅間よりも、新宿駅と大阪駅間の距離の誤差の方が大きい。
	// 2 つの点が離れているほど誤差が大きくなる。
	// 誤差を測る際、国土地理院のサイトを使用した。
	// https://vldb.gsi.go.jp/sokuchi/surveycalc/surveycalc/bl2stf.html
	d := l1.Distance(l2)
	fmt.Printf("新宿駅 %s -> 大阪駅 %s\n", l1, l2)
	fmt.Printf("距離 %f(%f km)\n", d, d*6371.01) // 誤差 500mぐらい
	d = l1.Distance(l3)
	fmt.Printf("新宿駅 %s -> 横浜駅 %s\n", l1, l3)
	fmt.Printf("距離 %f(%f km)\n", d, d*6371.01) // 誤差 50mぐらい
}
