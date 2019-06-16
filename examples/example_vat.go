package main

import (
	"fmt"
	"github.com/miguelbemartin/vat"
)

func main() {

	client := vat.NewClient()

	rate, err := client.Rate.Get("")
	if err != nil {
		// handle your error
	}

	fmt.Println("The rate for %d is %d", "", rate)
}
