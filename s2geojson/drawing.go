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
	os.Remove(filePath)
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return xerrors.Errorf("Cannot open file : %w", err)
	}
	defer f.Close()
	err = Fprint(f, points)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
	}
	return err
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
	_, err = fmt.Fprint(out, string(body))
	return err
}
