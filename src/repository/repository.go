package repository

import "gorm.io/gorm"

type Repository struct {
	Prodi     IProdiRepository
	Mahasiswa IMahasiswaRepository
}

func Init(db *gorm.DB) *Repository {
	return &Repository{
		Prodi: NewProdiRepository(db),
		Mahasiswa: NewMahasiswaRepository(db),
	}
}
