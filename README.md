# luhn

Simple [go](https://golang.org/) package that implements the
[Luhn algorithm](https://en.wikipedia.org/wiki/Luhn_algorithm).

## Installation

The latest and greatest version of this software is available at
[github.com/tcort/luhn](https://github.com/tcort/luhn).

    go get github.com/tcort/luhn

## API

For complete API documentation, see [godoc.org/github.com/tcort/luhn](https://godoc.org/github.com/tcort/luhn)

## Example Usage

Here's a simple command line tool that accepts a number as a
command line argument and reports whether or not it passed
the Luhn check:
```
package main

import (
    "fmt"
    "github.com/tcort/luhn"
    "os"
)

func main() {

    if len(os.Args) != 2 {
        fmt.Println("Usage: lc [number]")
        return
    }

    if luhn.Check(os.Args[1]) {
        fmt.Println("The number passed the Luhn check")
    } else {
        fmt.Println("The number failed the Luhn check")
    }
}
```

## License

See [LICENSE.md](https://github.com/tcort/luhn/blob/master/LICENSE.md)
