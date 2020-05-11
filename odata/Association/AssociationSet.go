package Association

import "encoding/xml"

type AssociationSet struct {
	XMLName xml.Name `xml:"AssociationSet"`

	Name        string `xml:"Name,attr"`
	Association string `xml:"Association,attr"`

	End                   []End                 `xml:"End"`
	ReferentialConstraint ReferentialConstraint `xml:"ReferentialConstraint"`
}
