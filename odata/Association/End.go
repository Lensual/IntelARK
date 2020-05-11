package Association

import "encoding/xml"

type End struct {
	XMLName xml.Name `xml:"End"`

	//Association
	Type         string `xml:"Type,attr"`
	Role         string `xml:"Role,attr"`
	Multiplicity string `xml:"Multiplicity,attr"`

	OnDelete OnDelete `xml:"OnDelete"`

	//AssociationSet
	EntitySet string `xml:"EntitySet,attr"`
}
