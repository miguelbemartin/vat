package main

import (
	"fmt"
	"github.com/miguelbemartin/vat"
)

func main() {

	exist, err := vat.Validate("XXXX")
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	fmt.Println("The result is ", exist)

	rate, err := vat.GetRate("ES")
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	fmt.Println("The result is ", rate)

}
