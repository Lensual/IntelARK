package ReferentialConstraint

import (
	"encoding/xml"

	. "github.com/Lensual/IntelARK/odata/PropertyRef"
)

type Dependent struct {
	XMLName xml.Name `xml:"Dependent"`

	Role string `xml:"Role,attr"`

	PropertyRef PropertyRef `xml:"PropertyRef"`
}
