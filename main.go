package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(rw, "Hello World")
		if err != nil {
			fmt.Println("Error occured ", err)
		}
		fmt.Printf("Byte written : %v\n", n)
	})

	http.ListenAndServe(":8000", nil)

}
