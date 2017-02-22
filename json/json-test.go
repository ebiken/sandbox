package main

import (
	"encoding/json"
	"fmt"
)

func main () {

	type Match_tbl_0 struct {
		//name string
		//value int
		Name string
		Value int
	}
	m := Match_tbl_0{ "namename", 11 }
	b, err := json.Marshal(m)
	fmt.Println(err)
	fmt.Printf("%s\n", b)
	fmt.Println(b)

}
