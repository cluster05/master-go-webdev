package main

import (
	"fmt"
	"net/http"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello home")
}

func About(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello About")
}

func main() {

	// errors.New("error string")

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(":8000", nil)

}
