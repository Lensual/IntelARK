package odata

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	. "github.com/Lensual/IntelARK/odata/EntityType"
)

type EntityType struct {
	XMLName xml.Name `xml:"EntityType"`

	Name string `xml:"Name,attr"`

	Key                Key                  `xml:"Key"`
	Property           []Property           `xml:"Property"`
	NavigationProperty []NavigationProperty `xml:"NavigationProperty"`
}

func QueryEntityType(url string) ([]EntityType, error) {
	res, err := Get(url, nil, nil)
	if err != nil {
		return nil, err
	}
	//Read
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//Unmarshal
	edxm := Edmx{}
	err = xml.Unmarshal(data, &edxm)
	if err != nil {
		return nil, err
	}
	return edxm.DataServices.Schema.EntityType, nil
}

func Get(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	//new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("new request is fail ")
	}
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	log.Printf("Go %s URL : %s \n", http.MethodGet, req.URL.String())
	return client.Do(req)
}
