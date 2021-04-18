package s2geojson

import (
	"fmt"
	"io"
	"os"

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

func Print(
	points *[]s2.Point,
) error {
	filePath := os.Getenv("FILE_PATH_GEOJSON")
	f, err := os.Open(filePath)
	if err != nil {
		return xerrors.Errorf("Cannot open file : %w", err)
	}
	return Fprint(f, points)
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
