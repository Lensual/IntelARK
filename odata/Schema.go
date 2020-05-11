package odata

import "encoding/xml"

type Schema struct {
	XMLName xml.Name `xml:"Schema"`

	Namespace string `xml:"Namespace,attr"`

	EntityType      []EntityType    `xml:"EntityType"`
	Association     []Association   `xml:"Association"`
	EntityContainer EntityContainer `xml:"EntityContainer"`
	// AssociationSet  []AssociationSet `xml:"AssociationSet"` //TODO
}
