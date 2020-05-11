package odata

import "encoding/xml"

type Edmx struct {
	Version string `xml:"Version,attr"`

	DataServices DataServices `xml:"DataServices"`
}
type DataServices struct {
	XMLName xml.Name `xml:"DataServices"`

	DataServiceVersion    string `xml:"DataServiceVersion,attr"`
	MaxDataServiceVersion string `xml:"MaxDataServiceVersion,attr"`

	Schema Schema `xml:"Schema"`
}
