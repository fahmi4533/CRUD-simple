package main

import (
	"net/http"

	passiencontroller "github.com/jeypc/crudterbaru/controller/passienController"
)

func main() {
	http.HandleFunc("/", passiencontroller.Index)
	http.HandleFunc("/pasien", passiencontroller.Index)
	http.HandleFunc("/pasien/index", passiencontroller.Index)
	http.HandleFunc("/pasien/add", passiencontroller.Add)
	http.HandleFunc("/pasien/edit", passiencontroller.Edit)
	http.HandleFunc("/pasien/delete", passiencontroller.Delete)

	http.ListenAndServe(":8080", nil)
}
