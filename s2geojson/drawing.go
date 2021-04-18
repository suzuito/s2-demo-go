package s2geojson

import (
	"fmt"
	"io"

	"github.com/golang/geo/s2"
	geojson "github.com/paulmach/go.geojson"
	"golang.org/x/xerrors"
)

func DrawAsGeoJSON(
	points *[]s2.Point,
) *geojson.FeatureCollection {
	fc := geojson.NewFeatureCollection()
	for _, p := range *points {
		g := NewGeometryFromPoint(&p)
		f := geojson.NewFeature(g)
		fc.Features = append(fc.Features, f)
	}
	return fc
}

func Fprint(
	out io.Writer,
	points *[]s2.Point,
) error {
	fc := DrawAsGeoJSON(points)
	body, err := fc.MarshalJSON()
	if err != nil {
		return xerrors.Errorf("Cannot marshal geo json : %w", err)
	}
	fmt.Fprint(out, string(body))
	return nil
}
