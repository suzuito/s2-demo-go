package entity

import (
	"encoding/xml"

	"golang.org/x/xerrors"
)

type XMLURLSet struct {
	XMLName           xml.Name `xml:"urlset"`
	URLs              []XMLURL `xml:"url"`
	XMLNSXsi          string   `xml:"xmlns:xsi,attr"`
	XMLNS             string   `xml:"xmlns,attr"`
	XsiSchemaLocation string   `xml:"xsi:schemaLocation,attr"`
}

func (x *XMLURLSet) Marshal() (string, error) {
	b, err := xml.MarshalIndent(x, "", "    ")
	if err != nil {
		return "", xerrors.Errorf("Cannot marshal xml : %w", err)
	}

	c := string(b)
	c = `<?xml version="1.0" encoding="UTF-8"?>` + "\n" + c
	return c, nil
}

type XMLURL struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}
