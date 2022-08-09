package taskingcontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/jeypc/go-crud/entities"
	"github.com/jeypc/go-crud/libraries"
	"github.com/jeypc/go-crud/models"
)

var Validation = libraries.NewValidation()
var TaskingModel = models.NewTaskingModel()

func Index(response http.ResponseWriter, request *http.Request) {

	tasking, _ := TaskingModel.FindAll()

	data := map[string]interface{}{
		"tasking": tasking,
	}

	temp, err := template.ParseFiles("views/tasking/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/tasking/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var tasking entities.Tasking
		tasking.Task = request.Form.Get("task")
		tasking.Assigne = request.Form.Get("assigne")
		tasking.DeadLine = request.Form.Get("dead_line")

		var data = make(map[string]interface{})

		vErrors := Validation.Struct(tasking)

		if vErrors != nil {
			data["tasking"] = tasking
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data Tasking Berhasil Disimpan"
			TaskingModel.Create(tasking)
		}

		temp, _ := template.ParseFiles("views/tasking/add.html")
		temp.Execute(response, data)

	}
}

func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var tasking entities.Tasking
		TaskingModel.Find(id, &tasking)
		data := map[string]interface{}{
			"tasking": tasking,
		}

		temp, err := template.ParseFiles("views/tasking/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var tasking entities.Tasking
		tasking.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		tasking.Task = request.Form.Get("task")
		tasking.Assigne = request.Form.Get("assigne")
		tasking.DeadLine = request.Form.Get("dead_line")

		var data = make(map[string]interface{})

		vErrors := Validation.Struct(tasking)

		if vErrors != nil {
			data["tasking"] = tasking
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data Tasking Berhasil Diperbaharui"
			TaskingModel.Update(tasking)
		}

		temp, _ := template.ParseFiles("views/tasking/edit.html")
		temp.Execute(response, data)

	}
}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	TaskingModel.Delete(id)

	http.Redirect(response, request, "/tasking", http.StatusSeeOther)
}

func IsDone(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	TaskingModel.IsDone(id)
	http.Redirect(response, request, "/tasking", http.StatusSeeOther)
}

// func UpdateDate(response http.ResponseWriter, request *http.Request) {
// 	queryString := request.URL.Query()
// 	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

// 	TaskingModel.UpdateDate(id)
// 	http.Redirect(response, request, "/tasking", http.StatusSeeOther)
// }
