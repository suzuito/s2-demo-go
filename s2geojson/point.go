package s2geojson

import (
	"github.com/golang/geo/s2"
	geojson "github.com/paulmach/go.geojson"
)

func NewGeometryFromPoint(p *s2.Point) *geojson.Geometry {
	l := s2.LatLngFromPoint(*p)
	return geojson.NewPointGeometry([]float64{
		l.Lng.Degrees(),
		l.Lat.Degrees(),
	})
}
