package service

import "final-project-pemin/src/repository"

type Service struct {
	Prodi     IProdiService
	Mahasiswa IMahasiswaService
}

func Init(repository *repository.Repository) *Service {
	return &Service{
		Prodi: NewProdiService(repository.Prodi),
		Mahasiswa: NewMahasiswaService(repository.Mahasiswa),
	}
}
