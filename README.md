# luhn

Simple [go](https://golang.org/) package that implements the
[Luhn algorithm](https://en.wikipedia.org/wiki/Luhn_algorithm).

## Installation

The latest and greatest version of this software is available at
[github.com/tcort/luhn](https://github.com/tcort/luhn).

    go get github.com/tcort/luhn

## API

### `luhn.Check(number)`

Accepts a string argument consisting of 2 or more digits and returns
`true` or `false` for `pass` or `fail`.

## Usage

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
