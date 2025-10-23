package main

import (
	"io"
	"net/http"
)

func genarateMessage(name string) string {
	if name == "" {
		name = "World"
	}

	///////////// Test for staticcheck /////////////////
	// unusedVar := "This will cause a staticcheck error"

	// if true == true {
	// 	fmt.Println("Always true")
	// }
	///////////////////////////////////////////////////

	return "Hello " + name + "\n"
}

func main() {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}
		message := genarateMessage(name)
		io.WriteString(w, message)
	}

	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
