package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "home.html")
}

func About(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "about.html")

}

func renderTemplate(rw http.ResponseWriter, templatefileName string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + templatefileName)
	err := parsedTemplate.Execute(rw, nil)
	if err != nil {
		fmt.Println("Error occured while rendering ", templatefileName)
		return
	}
}

func main() {

	// errors.New("error string")

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(":8000", nil)

}
