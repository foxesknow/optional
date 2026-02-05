[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)

# Optional
The Optional library provides functional-like optional support to Go.

### Usage

The default value for an optional value is `none`
```go
var data optional.Value[int]
```

You use the `Some` and `None` functions to initialize an optional value:
```go
package main

import (
	"fmt"
	"strconv"

	"github.com/foxesknow/optional"
)

func main() {
	number := StringToInt("10")

	if i, hasValue := number.Get(); hasValue {
		fmt.Printf("Got %v\n", i)
	}

	// Or...
	if n := StringToInt("20"); n.IsSome() {
		fmt.Printf("Got %v\n", n.MustGet())
	} else {
		fmt.Println("that wasn't a number!")
	}

	// When it's not a number
	if n := StringToInt("hello, world"); n.IsSome() {
		fmt.Printf("Got %v\n", n.MustGet())
	} else {
		fmt.Println("that wasn't a number!")
	}
}

func StringToInt(s string) optional.Value[int] {
	if i, err := strconv.Atoi(s); err == nil {
		return optional.Some(i)
	}

	return optional.None[int]()
}
```