package s2geojson

import (
	"github.com/golang/geo/s2"
	geojson "github.com/paulmach/go.geojson"
)

func NewGeometryFromLoop(l *s2.Loop) *geojson.Geometry {
	p := [][]float64{}
	n := l.NumVertices()
	for i := 0; i < n; i++ {
		v := l.Vertex(i)
		ll := s2.LatLngFromPoint(v)
		p = append(p, []float64{
			ll.Lng.Degrees(),
			ll.Lat.Degrees(),
		})
	}
	g := geojson.NewPolygonGeometry([][][]float64{p})
	return g
}

// NewS2LoopFromGeometry ...
func NewS2LoopFromGeometry(polygon [][]float64) *s2.Loop {
	points := []s2.Point{}
	for _, p := range polygon {
		a := s2.LatLngFromDegrees(p[1], p[0])
		point := s2.PointFromLatLng(a)
		points = append(points, point)
	}
	loop := s2.LoopFromPoints(points)
	return loop
}
