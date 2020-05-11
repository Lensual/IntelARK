package main

import (
	"fmt"

	"github.com/Lensual/IntelARK/odata"
)

func main() {
	et, err := odata.QueryEntityType("https://odata.intel.com/API/v1_0/Products/$metadata")
	if err != nil {
		fmt.Printf("%+v", err)
	} else {
		fmt.Printf("%+v", et)
		// b, err := json.Marshal(et)
		// if err != nil {
		// 	fmt.Printf("%+v", et)
		// 	return
		// }
		// var out bytes.Buffer
		// err = json.Indent(&out, b, "", "    ")
		// if err != nil {
		// 	fmt.Printf("%+v", et)
		// 	return
		// }
		// fmt.Println(out.String())
	}
}
