package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Lensual/IntelARK/odata"
)

func main() {
	edmx, err := odata.QueryEdmx("https://odata.intel.com/API/v1_0/Products/$metadata")
	if err != nil {
		fmt.Printf("%+v", err)
	} else {
		// fmt.Printf("%+v", edmx)
		str, _ := toJson(edmx)
		fmt.Printf("%+v", str)
	}
}

func toJson(obj interface{}) (string, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return "", err
	}
	return fmt.Sprint(out.String()), nil
}
