package EntityType

import "encoding/xml"

type Property struct {
	XMLName xml.Name `xml:"Property"`

	Name        string `xml:"Name,attr"`
	Type        string `xml:"Type,attr"`
	Nullable    bool   `xml:"Nullable,attr"`
	MaxLength   string `xml:"MaxLength,attr"`
	FixedLength string `xml:"FixedLength,attr"`
	Unicode     bool   `xml:"Unicode,attr"`
}
