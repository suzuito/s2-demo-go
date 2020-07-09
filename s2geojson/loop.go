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
