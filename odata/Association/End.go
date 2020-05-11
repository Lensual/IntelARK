package association

import "encoding/xml"

type End struct {
	XMLName xml.Name `xml:"End"`

	Type         string `xml:"Type,attr"`
	Role         string `xml:"Role,attr"`
	Multiplicity string `xml:"Multiplicity,attr"`

	// OnDelete OnDelete `xml:"OnDelete"`
}
