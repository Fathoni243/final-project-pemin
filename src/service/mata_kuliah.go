package service

import (
	"final-project-pemin/src/model"
	"final-project-pemin/src/repository"
)

type (
	IMataKuliahService interface {
		FindALl() ([]*model.MataKuliah, error)
	}

	mataKuliahService struct {
		mataKuliahRepository repository.IMataKuliahRepository
	}
)

func NewMataKuliahService(mataKuliahRepository repository.IMataKuliahRepository) IMataKuliahService {
	return &mataKuliahService{mataKuliahRepository: mataKuliahRepository}
}

func (mks *mataKuliahService) FindALl() ([]*model.MataKuliah, error) {
	matakuliahs, err := mks.mataKuliahRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return matakuliahs, nil
}