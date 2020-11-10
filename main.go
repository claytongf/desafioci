package main

import (
	"fmt"
	"math"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var x float64 = 0.0001

	for i := 0; i < 1000000; i++ {
		x += math.Sqrt(x)
	}

	fmt.Fprint(w, greeting("Code.education Rocks!"))
}

func greeting(str string) string {
	return "<b>" + str + "</b>"
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
