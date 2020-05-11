package Entity

import (
	"encoding/xml"

	. "github.com/Lensual/IntelARK/odata/Entity/EntityType"
)

type EntityType struct {
	XMLName xml.Name `xml:"EntityType"`

	Name string `xml:"Name,attr"`

	Key                Key                  `xml:"Key"`
	Property           []Property           `xml:"Property"`
	NavigationProperty []NavigationProperty `xml:"NavigationProperty"`
}
