package main

import (
	"fmt"
	"net/http"
)

const PORT string = ":8000"

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

	})

	fmt.Println("Server is runnig on PORT ", PORT)
	http.ListenAndServe(PORT, nil)
}

// nodemon --exec go run main.go --signal SIGTERM
