package ReferentialConstraint

import (
	"encoding/xml"

	. "github.com/Lensual/IntelARK/odata/PropertyRef"
)

type Principal struct {
	XMLName xml.Name `xml:"Principal"`

	Role string `xml:"Role,attr"`

	PropertyRef PropertyRef `xml:"PropertyRef"`
}
