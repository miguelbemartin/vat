# vat
vat is a [Go](https://golang.org/) client library for accessing the Open Service by [European Commission](https://ec.europa.eu/commission/index_en)

## Getting Started

### Prerequisites
- Go 1.11+

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

func main () {

    //...

    client := vat.NewClient()
    
    rate, err := client.Rate.Get("")
    if err != nil {
      // handle your error
    }
    
    fmt.Println("The rate for %d is %d", "", rate)

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
