# sensible-terminal

[![Documentation](https://godoc.org/github.com/particleflux/sensible-terminal?status.svg)](http://godoc.org/github.com/particleflux/sensible-terminal)
[![CircleCI](https://circleci.com/gh/particleflux/sensible-terminal.svg?style=shield)](https://circleci.com/gh/particleflux/sensible-terminal)
[![Go Report Card](https://goreportcard.com/badge/github.com/particleflux/sensible-terminal)](https://goreportcard.com/report/github.com/particleflux/sensible-terminal)

Auto-detect installed terminal emulators

## What

This will search through a builtin list of known terminal emulators, like
`xterm` or `gnome-terminal`. Those binaries which exist in PATH are returned.

Inspired by [i3-sensible-terminal](https://github.com/i3/i3/blob/next/i3-sensible-terminal)

## Example

```go
package main

import "fmt"
import "github.com/particleflux/sensible-terminal"

func main() {
    term, err := sensibleterminal.First()
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("first terminal:", term)


    terms, _ := sensibleterminal.All()
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
