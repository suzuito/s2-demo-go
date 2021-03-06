package sgin

import (
	"github.com/golang/geo/s2"
	geojson "github.com/paulmach/go.geojson"
	"github.com/suzuito/s2-demo-go/s2geojson"
)

// LatLngLiteral ...
type LatLngLiteral struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// NewS2LoopFromLatLngLiteralArray ...
func NewS2LoopFromLatLngLiteralArray(a []LatLngLiteral) *s2.Loop {
	points := []s2.Point{}
	for _, v := range a {
		points = append(points, s2.PointFromLatLng(
			s2.LatLngFromDegrees(
				v.Lat,
				v.Lng,
			),
		))
	}
	return s2.LoopFromPoints(points)
}

// NewLatLngLiteralArrayFromS2Loop ...
func NewLatLngLiteralArrayFromS2Loop(l *s2.Loop) *[]LatLngLiteral {
	ret := []LatLngLiteral{}
	for _, v := range l.Vertices() {
		latlng := s2.LatLngFromPoint(v)
		ret = append(ret, LatLngLiteral{
			Lat: latlng.Lat.Degrees(),
			Lng: latlng.Lng.Degrees(),
		})
	}
	return &ret
}

// CellLiteral ...
type CellLiteral struct {
	ID         string           `json:"id"`
	GeoJSON    *geojson.Feature `json:"geoJson"`
	Center     LatLngLiteral    `json:"center"`
	Level      int              `json:"level"`
	ApproxArea float64          `json:"approxArea"`
	Points     []s2.Point       `json:"points"`
}

// NewCellLiteralFromCell ...
func NewCellLiteralFromCell(c s2.Cell) *CellLiteral {
	cid := c.ID()
	center := s2.LatLngFromPoint(c.Center())
	points := s2.LoopFromCell(c).Vertices()
	return &CellLiteral{
		ID: cid.ToToken(),
		GeoJSON: geojson.NewFeature(
			s2geojson.NewGeometryFromLoop(s2.LoopFromCell(c)),
		),
		Center: LatLngLiteral{
			Lat: center.Lat.Degrees(),
			Lng: center.Lng.Degrees(),
		},
		Level:      c.Level(),
		ApproxArea: c.ApproxArea(),
		Points:     points,
	}
}

// NewCellLiteralsFromCellIDs ...
func NewCellLiteralsFromCellIDs(cellIDs []s2.CellID) *[]*CellLiteral {
	res := []*CellLiteral{}
	for _, cellID := range cellIDs {
		res = append(res, NewCellLiteralFromCell(s2.CellFromCellID(cellID)))
	}
	return &res
}
