package main

import (
	"fmt"
	"testing"
)

func TestGreeting(t *testing.T) {
	got := greeting("Code.education Rocks!")
	want := "<b>Code.education Rocks!</b>"

	if got != want {
		t.Errorf("greeting() \n got: %v \n want: %v", got, want)
	} else {
		fmt.Println(greeting("Code.education Rocks!"))
	}
}
