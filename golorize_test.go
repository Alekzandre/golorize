package golorize

import (
	"testing"
)

var test1 = []string{}
var test2 = []string{"first"}
var test3 = []string{"first", "second"}
var test4 = []string{"first", "second", "third"}
var test5 = []string{"first", "second", "third", "fourth"}

func TestOptionSize(t *testing.T) {
	res1, err1 := OptionsSize(test1)
	if res1 != 0 && err1 != nil {
		t.Error("expected res = 0 et err == nil")
	}
	res2, err2 := OptionsSize(test2)
	if res2 != 1 && err2 != nil {
		t.Error("expected res = 1 et err == nil")
	}
	res3, err3 := OptionsSize(test3)
	if res3 != 2 && err3 != nil {
		t.Error("expected res = 2 et err == nil")
	}
	res4, err4 := OptionsSize(test4)
	if res4 != 3 && err4 != nil {
		t.Error("expected res = 3 et err == nil")
	}
	res5, err5 := OptionsSize(test5)
	if res5 != 0 && err5 == nil {
		t.Error("expected res = 0 et err != nil")
	}
}

func TestNewGolorize(t *testing.T) {
	Golorize := NewGolorize()
	if Golorize == nil {
		t.Error("expected the golorize object to be created")
	}
}
