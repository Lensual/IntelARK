package EntityType

import (
	"encoding/xml"

	. "github.com/Lensual/IntelARK/odata/PropertyRef"
)

type Key struct {
	XMLName xml.Name `xml:"Key"`

	PropertyRef []PropertyRef `xml:"PropertyRef"`
}
