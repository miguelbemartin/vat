package main

import (
	"fmt"
	"github.com/miguelbemartin/vat"
)

func main() {

	client := vat.NewClient()

	exist, err := client.Validate("ESXX")
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	fmt.Println("The result is ", exist)

	rate, err := client.GetRate("ES")
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	fmt.Println("The result is ", rate)

}
