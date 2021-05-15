package main

import (
	"fmt"

	"github.com/golang/geo/s2"
)

func main() {
	a := s2.PointFromCoords(1, 0, 0)
	b := s2.PointFromCoords(0, 0, 0)
	c := s2.PointFromCoords(-1, 0, 0)
	ang := s2.Angle(a, b, c)
	fmt.Printf(
		"%.2f, %.2f -> %.2f, %.2f -> %.2f, %.2f によってできる角の角度は %.2f %.2f度\n",
		a.X, a.Y,
		b.X, b.Y,
		c.X, c.Y,
		ang,
		ang.Degrees(),
	)

	a = s2.PointFromCoords(1, 0, 0)
	b = s2.PointFromCoords(0, 0, 0)
	c = s2.PointFromCoords(0, 1, 0)
	ang = s2.Angle(a, b, c)
	fmt.Printf(
		"%.2f, %.2f -> %.2f, %.2f -> %.2f, %.2f によってできる角の角度は %.2f %.2f度\n",
		a.X, a.Y,
		b.X, b.Y,
		c.X, c.Y,
		ang,
		ang.Degrees(),
	)

	a = s2.PointFromCoords(1, 0, 0)
	b = s2.PointFromCoords(0, 0, 0)
	c = s2.PointFromCoords(1, 1, 0)
	ang = s2.Angle(a, b, c)
	fmt.Printf(
		"%.2f, %.2f -> %.2f, %.2f -> %.2f, %.2f によってできる角の角度は %.2f %.2f度\n",
		a.X, a.Y,
		b.X, b.Y,
		c.X, c.Y,
		ang,
		ang.Degrees(),
	)
}
