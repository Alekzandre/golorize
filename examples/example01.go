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
