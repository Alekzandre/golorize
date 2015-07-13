# golorize
*Golorize is a small package that makes your go experience more colorfull.*

It's in [Google's Go](https://golang.org/) language to colorize your syntax.

####To install the package make sure your $GOPATH is correctly set.
```
go get "github.com/Alekzandre/golorize"
```
####Exemple:
```
package main

import (
	glz "github.com/Alekzandre/golorize"
)

func main() {
	gz := glz.NewGolorize()
	gz.Color("basic on Stdout")
	gz.Color("basic on Stderr", "2")
	gz.Color("underlined", "1", glz.UNDERLINED)
	gz.Color("bold and yellow", "1", glz.BOLD, glz.F_YELLOW)
	gz.Color("reversed so red on cyan background", "1", glz.REVERSE, glz.F_CYAN, glz.B_RED)
	gz.Color("It makes errors more flashy", "2", glz.BOLD, glz.F_RED, glz.B_BLUE)
}

```
![out_example01](/examples/example01.png)

####Features to come:
Improve the code.

Make it easier to use.

