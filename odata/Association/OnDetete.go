package Association

import "encoding/xml"

type OnDelete struct {
	XMLName xml.Name `xml:"OnDelete"`

	Action string `xml:"Action,attr"`
}
