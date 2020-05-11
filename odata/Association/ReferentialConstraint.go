package association

import (
	"encoding/xml"

	. "github.com/Lensual/IntelARK/odata/Association/ReferentialConstraint"
)

type ReferentialConstraint struct {
	XMLName xml.Name `xml:"ReferentialConstraint"`

	Principal Principal `xml:"Principal"`
	Dependent Dependent `xml:"Dependent"`
}
