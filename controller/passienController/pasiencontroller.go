package passiencontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/jeypc/crudterbaru/entities"
	"github.com/jeypc/crudterbaru/libraries"
	"github.com/jeypc/crudterbaru/models"
)

var validation = libraries.NewValidator()
var pasienModels = models.NewPasien()

func Index(w http.ResponseWriter, r *http.Request) {

	pasien, _ := pasienModels.FindAll()

	data := map[string]interface{}{
		"pasien": pasien,
	}

	temp, err := template.ParseFiles("views/pasien/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}
func Add(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		temp, err := template.ParseFiles("views/pasien/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {

		r.ParseForm()
		var pasien entities.Pasien
		pasien.NamaLengkap = r.Form.Get("nama_lengkap")
		pasien.NIK = r.Form.Get("nik")
		pasien.JenisKelamin = r.Form.Get("jenis_kelamin")
		pasien.TemapatLahir = r.Form.Get("tempat_lahir")
		pasien.TanggalLahir = r.Form.Get("tanggal_lahir")
		pasien.Alamat = r.Form.Get("alamat")
		pasien.NoHp = r.Form.Get("no_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(pasien)
		if vErrors != nil {
			data["pasien"] = pasien
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data berhasil disimpan"
			pasienModels.Create(pasien)
		}

		temp, _ := template.ParseFiles("views/pasien/add.html")
		temp.Execute(w, data)

	}

}
func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		queryStrings := r.URL.Query()
		id, _ := strconv.ParseInt(queryStrings.Get("id"), 10, 64)

		var pasien entities.Pasien
		pasienModels.Find(id, &pasien)

		data := map[string]interface{}{
			"pasien": pasien,
		}

		temp, err := template.ParseFiles("views/pasien/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, data)
	} else if r.Method == http.MethodPost {

		r.ParseForm()
		var pasien entities.Pasien
		pasien.Id, _ = strconv.ParseInt(r.Form.Get("id"), 10, 64)
		pasien.NamaLengkap = r.Form.Get("nama_lengkap")
		pasien.NIK = r.Form.Get("nik")
		pasien.JenisKelamin = r.Form.Get("jenis_kelamin")
		pasien.TemapatLahir = r.Form.Get("tempat_lahir")
		pasien.TanggalLahir = r.Form.Get("tanggal_lahir")
		pasien.Alamat = r.Form.Get("alamat")
		pasien.NoHp = r.Form.Get("no_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(pasien)
		if vErrors != nil {
			data["pasien"] = pasien
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data berhasil diperbaruhi"
			pasienModels.Update(pasien)
		}

		temp, _ := template.ParseFiles("views/pasien/edit.html")
		temp.Execute(w, data)

	}
}
func Delete(w http.ResponseWriter, r *http.Request) {

	queryStrings := r.URL.Query()
	id, _ := strconv.ParseInt(queryStrings.Get("id"), 10, 64)

	pasienModels.Delete(id)

	http.Redirect(w, r, "/pasien", http.StatusSeeOther)
}
func main()
