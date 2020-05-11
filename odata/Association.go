package odata

import (
	"encoding/xml"

	. "github.com/Lensual/IntelARK/odata/Association"
)

type Association struct {
	XMLName xml.Name `xml:"Association"`

	Name string `xml:"Name,attr"`

	End                   []End                 `xml:"End"`
	ReferentialConstraint ReferentialConstraint `xml:"ReferentialConstraint"`
}
