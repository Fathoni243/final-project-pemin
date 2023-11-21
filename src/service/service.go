package service

import "final-project-pemin/src/repository"

type Service struct {
	Prodi IProdiService
}

func Init(repository *repository.Repository) *Service {
	return &Service{
		Prodi: NewProdiService(repository.Prodi),
	}
}