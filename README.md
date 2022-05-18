# gocolor

Goでコンソールに出力する文字色を変えれます。

# How to use

## Install

```bash
go get github.com/code-raisan/gocolor@latest
```

## Sample

```go
package main

import (
	"github.com/code-raisan/gocolor"
)

func main() {
    // Out put is red
    println(gocolor.Red("gocolor.Red()"))
	
    // Out put is blue
    println(gocolor.Blue("gocolor.Blue()"))
	
    // Out put is yellow
    println(gocolor.Yellow("gocolor.Yellow()"))
	
    // Out put is green
    println(gocolor.Green("gocolor.Green()"))
	
    // Out put is magenta
    println(gocolor.Purple("gocolor.Magenta()"))
	
    // Out put is cayn
    println(gocolor.Cayn("gocolor.Cayn()"))
	
    // Out put is white
    println(gocolor.White("gocolor.White()"))
	
    // Out put is black
    println(gocolor.Black("gocolor.Black()"))
}

```