package service

import (
	"final-project-pemin/src/model"
	"final-project-pemin/src/repository"
)

type (
	IProdiService interface {
		FindAll() ([]*model.Prodi, error)
	}

	prodiService struct {
		prodiRepository repository.IProdiRepository
	}
)

func NewProdiService(prodiRepository repository.IProdiRepository) IProdiService {
	return &prodiService{prodiRepository: prodiRepository}
}

func (ps *prodiService) FindAll() ([]*model.Prodi, error) {
	prodis, err := ps.prodiRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return prodis, nil
}