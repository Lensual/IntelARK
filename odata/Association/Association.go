package Association

import (
	"encoding/xml"
)

type Association struct {
	XMLName xml.Name `xml:"Association"`

	Name string `xml:"Name,attr"`

	End                   []End                 `xml:"End"`
	ReferentialConstraint ReferentialConstraint `xml:"ReferentialConstraint"`
}
