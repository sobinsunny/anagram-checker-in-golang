package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("input_data.txt")
	if err != nil {
		fmt.Println(err)
	}
	var data string = string(b)
	fmt.Println(data)

}
