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
