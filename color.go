package golorize

import (
	"bytes"
	"fmt"
	"os"
)

func (g Golorize) Color(s string, options ...string) {

	var buffer bytes.Buffer
	opNum, err := OptionsSize(options)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if opNum == 0 {
		buffer.WriteString(START)
		buffer.WriteString(NORMAL)
		buffer.WriteString(F_WHITE)
		buffer.WriteString(LIMIT)
		buffer.WriteString(s)
		buffer.WriteString(END)
		fmt.Println(buffer.String())
	} else if opNum == 1 {
		buffer.WriteString(START)
		buffer.WriteString(string(options[0]))
		buffer.WriteString(F_WHITE)
		buffer.WriteString(LIMIT)
		buffer.WriteString(s)
		buffer.WriteString(END)
		fmt.Println(buffer.String())
	} else if opNum == 2 {
		buffer.WriteString(START)
		buffer.WriteString(string(options[0]))
		buffer.WriteString(options[1])
		buffer.WriteString(LIMIT)
		buffer.WriteString(s)
		buffer.WriteString(END)
		fmt.Println(buffer.String())
	} else {
		buffer.WriteString(START)
		buffer.WriteString(string(options[0]))
		buffer.WriteString(options[1])
		buffer.WriteString(options[2])
		buffer.WriteString(LIMIT)
		buffer.WriteString(s)
		buffer.WriteString(END)
		fmt.Println(buffer.String())
	}
}
