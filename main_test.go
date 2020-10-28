package main

import (
	"fmt"
	"testing"
)

func TestGreeting(t *testing.T) {
	got := greeting("Code.education Rocks!")
	want := "Code.education Rocks!"

	if got != want {
		t.Errorf("greeting() \n got: %v \n want: %v", got, want)
	} else {
		fmt.Println(greeting("Code.education Rocks!"))
	}
}
