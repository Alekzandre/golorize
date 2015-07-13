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
	gz.Color("basic")
	gz.Color("underlined", glz.UNDERLINED)
	gz.Color("bold and yellow", glz.BOLD, glz.F_YELLOW)
	gz.Color("reversed so red on cyan background", glz.REVERSE, glz.F_CYAN, glz.B_RED)
	gz.Color("It makes errors more flashy", glz.BOLD, glz.F_RED, glz.B_BLUE)
}
```
####Features to come:
Choose between Stderr and Stdout.
Improve the code.
Makes it easier to use.

