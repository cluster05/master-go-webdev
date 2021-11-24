package api

import (
	"fmt"
	"net/http"
)

var tasks []Task = make([]Task, 0)

func createTask(rw http.ResponseWriter, r *http.Request) {

	fmt.Printf("%v", r)

}

func readTask(rw http.ResponseWriter, r *http.Request) {

}

func readTasks(rw http.ResponseWriter, r *http.Request) {

}

func updateTask(rw http.ResponseWriter, r *http.Request) {

}

func deleteTask(rw http.ResponseWriter, r *http.Request) {

}
