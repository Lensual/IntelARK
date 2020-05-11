package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Lensual/IntelARK/odata"
	"github.com/jinzhu/gorm"

	"io/ioutil"

	"encoding/base64"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type public_mobile_data struct {
	D D `json:"d"`
}

type D struct {
	Language string `json:"language"`
	Checksum string `json:"checksum"`
	DataBin  string `json:"data_bin"`
}

func main() {
	language := "zh-cn"
	timestamp := fmt.Sprintf("%v", time.Now().UnixNano())

	res, err := Get("https://odata.intel.com/API/v1_0/Products/public_mobile_data('"+language+"')?"+
		"format=json&"+timestamp, nil, nil)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	pmd := public_mobile_data{}
	err = json.Unmarshal(data, &pmd)
	if err != nil {
		return
	}

	bin, err := base64.StdEncoding.DecodeString(pmd.D.DataBin)
	if err != nil {
		return
	}

	filename := "./" + pmd.D.Checksum + ".sqlite"

	err = ioutil.WriteFile(filename, bin, os.ModePerm)
	if err != nil {
		return
	}

	db, err := gorm.Open("sqlite3", filename)
	defer db.Close()
	if err != nil {
		return
	}
}

func edmx() {
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
