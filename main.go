package main

import (
	"net/http"

	"github.com/jeypc/go-crud/controllers/taskingcontroller"
)

func main() {

	//routing

	http.HandleFunc("/", taskingcontroller.Index)
	http.HandleFunc("/tasking", taskingcontroller.Index)
	http.HandleFunc("/tasking/index", taskingcontroller.Index)
	http.HandleFunc("/tasking/add", taskingcontroller.Add)
	http.HandleFunc("/tasking/edit", taskingcontroller.Edit)
	http.HandleFunc("/tasking/delete", taskingcontroller.Delete)
	http.HandleFunc("/tasking/isdone", taskingcontroller.IsDone)

	//listen
	http.ListenAndServe(":3000", nil)
}
