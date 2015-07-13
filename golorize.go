package golorize

import (
	"errors"
)

// Basic attributes

const (
	START = "\033["
	END   = "\033[0m"
	LIMIT = "m"
)

// Basic streams (default: Stdout)

// Basic options

const (
	NORMAL     = "0;"
	BOLD       = "1;"
	UNDERLINED = "4;"
	BLINKING   = "5;"
	REVERSE    = "7;"
)

// Basic foreground color

const (
	F_BLACK   = "30"
	F_RED     = "31"
	F_GREEN   = "32"
	F_YELLOW  = "33"
	F_BLUE    = "34"
	F_MAGENTA = "35"
	F_CYAN    = "36"
	F_WHITE   = "37"
)

// Basic background color

const (
	B_BLACK   = ";40"
	B_RED     = ";41"
	B_GREEN   = ";42"
	B_YELLOW  = ";43"
	B_BLUE    = ";44"
	B_MAGENTA = ";45"
	B_CYAN    = ";46"
	B_WHITE   = ";47"
)

type Golorize struct{}

// Calculate the number of options must be in range 0-3

func OptionsSize(options []string) (int, error) {
	if len(options) > 3 {
		err := errors.New("Too many arguments.")
		return 0, err
	}
	return len(options), nil
}

func NewGolorize() *Golorize {
	return &Golorize{}
}
