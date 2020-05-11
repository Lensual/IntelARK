package odata

import (
	"encoding/xml"
	"io/ioutil"
)

func QueryEdmx(url string) (Edmx, error) {
	edmx := Edmx{}
	res, err := Get(url, nil, nil)
	if err != nil {
		return edmx, err
	}
	//Read
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return edmx, err
	}
	//Unmarshal
	err = xml.Unmarshal(data, &edmx)
	if err != nil {
		return edmx, err
	}
	return edmx, nil
}
