package main

import (
	"fmt"
)

func main() {

	var aValue int

	defer fmt.Println(aValue)

	for i := 0; i < 100; i++ {
		aValue++
	}

}
