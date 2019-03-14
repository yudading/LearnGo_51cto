package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "a.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}
}
