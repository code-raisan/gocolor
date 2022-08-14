# gocolor

![](https://img.shields.io/github/stars/code-raisan/gocolor?style=flat-square)

Goでコンソールに出力する文字色を変えれます。

Change the color of characters output to the console with Go.

# How to use

## Install

```bash
go get github.com/code-raisan/gocolor@latest
```

## Sample

`gocolor.Default()` を使用すると通常の文字色に戻すことができます

### 色文字の出力例

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
	
    // Out put is purple(magenta)
    println(gocolor.Purple("gocolor.Purple()"))
	
    // Out put is cyan
    println(gocolor.Cyan("gocolor.Cyan()"))
	
    // Out put is white
    println(gocolor.White("gocolor.White()"))
	
    // Out put is black
    println(gocolor.Black("gocolor.Black()"))
}

```

![image](https://user-images.githubusercontent.com/67790884/169005087-cea1fecf-be46-47db-a2cb-8891b60c69b6.png)

