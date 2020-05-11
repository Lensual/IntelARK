package Entity

import "encoding/xml"

type EntitySet struct {
	XMLName xml.Name `xml:"EntitySet"`

	Name       string `xml:"Name,attr"`
	EntityType string `xml:"EntityType,attr"`
}
