package s2geojson

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/geo/s2"
	geojson "github.com/paulmach/go.geojson"
	"golang.org/x/xerrors"
)

type PrintGeoJSONOptionLatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type PrintGeoJSONOption struct {
	StyleHeight string                   `json:"styleHeight"`
	Zoom        int                      `json:"zoom"`
	Center      PrintGeoJSONOptionLatLng `json:"center"`
}

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
	option *PrintGeoJSONOption,
) error {
	dirPathResult := os.Getenv("DIR_PATH_RESULT")
	filePathGeoJSON := filepath.Join(dirPathResult, "result.geojson")
	filePathGeoJSONOption := filepath.Join(dirPathResult, "result.geojson.option.json")
	os.Remove(filePathGeoJSON)
	os.Remove(filePathGeoJSONOption)
	bytesGeoJSON, _ := DrawAsGeoJSON(points).MarshalJSON()
	bytesGeoJSONOption, _ := json.MarshalIndent(option, "", " ")
	if err := ioutil.WriteFile(filePathGeoJSON, bytesGeoJSON, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
	}
	if err := ioutil.WriteFile(filePathGeoJSONOption, bytesGeoJSONOption, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
	}
	return nil
}

func Fprint(
	outGeoJSON io.Writer,
	points *[]s2.Point,
	option *PrintGeoJSONOption,
) error {
	fc := DrawAsGeoJSON(points)
	body, err := fc.MarshalJSON()
	if err != nil {
		return xerrors.Errorf("Cannot marshal geo json : %w", err)
	}
	_, err = fmt.Fprint(outGeoJSON, string(body))
	return err
}
