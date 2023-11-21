package model

type (
	Mahasiswa struct {
		NIM        string       `json:"nim" gorm:"primaryKey;type:varchar(100);unique"`
		ProdiId    int64        `json:"id_prodi" gorm:"not null"`
		Nama       string       `json:"nama" gorm:"not null;type:varchar(100)"`
		Angkatan   int32        `json:"angkatan" gorm:"not null"`
		Password   string       `json:"password" gorm:"not null;type:varchar(100)"`
		Prodi      Prodi        `json:"prodi"`
		MataKuliah []MataKuliah `json:"mata_kuliahs" gorm:"many2many:mahasiswa_matakuliah"`
	}

	MahasiswaInputRequest struct {
		NIM      string `json:"nim" binding:"required"`
		ProdiId  int64  `json:"id_prodi" binding:"required"`
		Nama     string `json:"nama" binding:"required"`
		Angkatan int32  `json:"angkatan" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	MahasiswaLoginRequest struct {
		NIM      string `json:"nim" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
)
