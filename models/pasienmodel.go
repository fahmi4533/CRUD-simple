package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jeypc/crudterbaru/config"
	"github.com/jeypc/crudterbaru/entities"
)

type PasienModels struct {
	con *sql.DB
}

func NewPasien() *PasienModels {
	con, err := config.DbConnection()
	if err != nil {
		panic(err)
	}
	return &PasienModels{
		con: con,
	}
}
func (p *PasienModels) FindAll() ([]entities.Pasien, error) {
	rows, err := p.con.Query("select * from pasien")
	if err != nil {
		return []entities.Pasien{}, err
	}
	defer rows.Close()

	var dataPas []entities.Pasien
	for rows.Next() {
		var pasien entities.Pasien
		rows.Scan(&pasien.Id,
			&pasien.NamaLengkap,
			&pasien.NIK,
			&pasien.JenisKelamin,
			&pasien.TemapatLahir,
			&pasien.TanggalLahir,
			&pasien.Alamat,
			&pasien.NoHp)

		if pasien.JenisKelamin == "1" {
			pasien.JenisKelamin = "Laki-laki 🥲"
		} else {
			pasien.JenisKelamin = "Wanita masjid"
		}
		tgl_lahir, _ := time.Parse("2006-01-02", pasien.TanggalLahir)

		pasien.TanggalLahir = tgl_lahir.Format("02-01-2006")

		dataPas = append(dataPas, pasien)
	}
	return dataPas, nil
}

func (p *PasienModels) Create(pasien entities.Pasien) bool {
	res, err := p.con.Exec("insert into pasien (nama_lengkap, nik, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat, no_hp) values(?, ?, ?, ?, ?, ?, ?)",
		pasien.NamaLengkap, pasien.NIK, pasien.JenisKelamin, pasien.TemapatLahir, pasien.TanggalLahir, pasien.Alamat, pasien.NoHp)

	if err != nil {
		fmt.Println(err)
		return false
	}
	lastId, _ := res.LastInsertId()

	return lastId > 0
}

func (p *PasienModels) Find(id int64, pasien *entities.Pasien) error {
	return p.con.QueryRow("select * from pasien where id = ?", id).Scan(
		&pasien.Id,
		&pasien.NamaLengkap,
		&pasien.NIK,
		&pasien.JenisKelamin,
		&pasien.TemapatLahir,
		&pasien.TanggalLahir,
		&pasien.Alamat,
		&pasien.NoHp)
}

func (p *PasienModels) Update(pasien entities.Pasien) error {

	_, err := p.con.Exec("update pasien set nama_lengkap=?, nik=?, jenis_kelamin=?, tempat_lahir=?, tanggal_lahir=?, alamat=?, no_hp =? where id = ?",
		pasien.NamaLengkap, pasien.NIK, pasien.JenisKelamin, pasien.TemapatLahir, pasien.TanggalLahir, pasien.Alamat, pasien.NoHp, pasien.Id)

	if err != nil {
		return err
	}
	return nil
}

func (p *PasienModels) Delete(id int64) {
	p.con.Exec("delete from pasien where id =?", id)
}
