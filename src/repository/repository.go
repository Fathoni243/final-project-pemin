package repository

import "gorm.io/gorm"

type Repository struct {
	Prodi      IProdiRepository
	Mahasiswa  IMahasiswaRepository
	MataKuliah IMataKuliahRepository
}

func Init(db *gorm.DB) *Repository {
	return &Repository{
		Prodi:     NewProdiRepository(db),
		Mahasiswa: NewMahasiswaRepository(db),
		MataKuliah: NewMataKuliahRepository(db),
	}
}
