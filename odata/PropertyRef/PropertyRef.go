package PropertyRef

import "encoding/xml"

type PropertyRef struct {
	XMLName xml.Name `xml:"PropertyRef"`

	Name string `xml:"Name,attr"`
}
