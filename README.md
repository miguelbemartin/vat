# vat
vat is a [Go](https://golang.org/) client library for accessing the Open Service by [European Commission](https://ec.europa.eu/commission/index_en)

[![Go Report Card](https://goreportcard.com/badge/github.com/miguelbemartin/vat)](https://goreportcard.com/report/github.com/miguelbemartin/vat)
[![Go](https://github.com/miguelbemartin/vat/actions/workflows/go.yml/badge.svg)](https://github.com/miguelbemartin/vat/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/miguelbemartin/vat/branch/master/graph/badge.svg)](https://codecov.io/gh/miguelbemartin/vat)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/miguelbemartin/vat/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://pkg.go.dev/github.com/miguelbemartin/vat)

## Getting Started

### Prerequisites
- Go 1.13+

### Installation
Run into the terminal the next command

```
go get github.com/miguelbemartin/vat
```

## Usage
```go
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
```

## Run the tests
```
go test . -v
```

## Contributing

## Authors
* **Miguel Ángel Martín** - [@miguelbemartin](https://twitter.com/miguelbemartin)

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details



https://ec.europa.eu/taxation_customs/sites/taxation/files/resources/documents/taxation/vat/how_vat_works/rates/vat_rates_en.pdf

