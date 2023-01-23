package entities

type Pasien struct {
	Id           int64
	NamaLengkap  string `validate:"required"`
	NIK          string `validate:"required"`
	JenisKelamin string `validate:"required"`
	TemapatLahir string `validate:"required"`
	TanggalLahir string `validate:"required"`
	Alamat       string `validate:"required"`
	NoHp         string `validate:"required"`
}
