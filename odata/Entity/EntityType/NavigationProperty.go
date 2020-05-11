package EntityType

import "encoding/xml"

type NavigationProperty struct {
	XMLName xml.Name `xml:"NavigationProperty"`

	Name         string `xml:"Name,attr"`
	Relationship string `xml:"Relationship,attr"`
	ToRole       string `xml:"ToRole,attr"`
	FromRole     string `xml:"FromRole,attr"`
}
