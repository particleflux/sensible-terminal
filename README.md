# sensible-terminal

[![Documentation](https://godoc.org/github.com/particleflux/sensible-terminal?status.svg)](http://godoc.org/github.com/particleflux/sensible-terminal)

Auto-detect installed terminal emulators

## What

This will search through a builtin list of known terminal emulators, like
`xterm` or `gnome-terminal`. Those binaries which exist in PATH are returned.

## Example

```go
package main

import "fmt"
import "github.com/particleflux/sensible-terminal"

func main() {
    term, err := sensible_terminal.First()
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("first terminal:", term)


    terms, _ := sensible_terminal.All()
    if err != nil {
        panic(err.Error())
    }

    fmt.Println("all terminals:")
    for _, term := range terms {
        fmt.Println(term)
    }
}
```

Will print on my system:

```text
first terminal: urxvt
all terminals:
urxvt
xterm
```
