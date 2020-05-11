package odata

import "encoding/xml"

type EntityContainer struct {
	XMLName xml.Name `xml:"EntityContainer"`

	Name                     string `xml:"Name,attr"`
	IsDefaultEntityContainer bool   `xml:"IsDefaultEntityContainer,attr"`

	EntitySet []EntitySet `xml:"EntitySet"`
}
