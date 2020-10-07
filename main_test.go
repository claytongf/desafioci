package main

import ("testing"; "fmt")

func TestSoma(t *testing.T){
	got := soma(5, 5)
	want := 10

	if got != want {
		t.Errorf("soma(5, 5) \n got: %v \n want: %v", got, want)
	}else{
		fmt.Println(soma(5, 5))
	}
}