package model

type (
	Mahasiswa struct {
		NIM        string       `json:"nim" gorm:"primaryKey;type:varchar(15);unique"`
		ProdiId    int64        `json:"prodiId" gorm:"not null"`
		Nama       string       `json:"nama" gorm:"not null;type:varchar(100)"`
		Angkatan   int32        `json:"angkatan" gorm:"not null"`
		Password   string       `json:"password" gorm:"not null;type:varchar(100)"`
		Prodi      Prodi        `json:"prodi"`
		MataKuliah []MataKuliah `json:"matakuliah" gorm:"many2many:mahasiswa_matakuliah"`
	}

	MahasiswaInputRequest struct {
		NIM      string `json:"nim" binding:"required"`
		ProdiId  string `json:"prodiId" binding:"required"`
		Nama     string `json:"nama" binding:"required"`
		Angkatan string `json:"angkatan" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	MahasiswaLoginRequest struct {
		NIM      string `json:"nim" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
)
