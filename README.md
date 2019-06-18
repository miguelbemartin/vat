# vat
vat is a [Go](https://golang.org/) client library for accessing the Open Service by [European Commission](https://ec.europa.eu/commission/index_en)

[![Go Report Card](https://goreportcard.com/badge/github.com/miguelbemartin/vat)](https://goreportcard.com/report/github.com/miguelbemartin/vat)
[![Build Status](https://travis-ci.org/miguelbemartin/vat.svg?branch=master)](https://travis-ci.org/miguelbemartin/vat)
[![codecov](https://codecov.io/gh/miguelbemartin/vat/branch/master/graph/badge.svg)](https://codecov.io/gh/miguelbemartin/vat)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/miguelbemartin/vat/master/LICENSE)

## Getting Started

### Prerequisites
- Go 1.10+

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
