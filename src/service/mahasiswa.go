package service

import (
	"errors"
	"final-project-pemin/src/model"
	"final-project-pemin/src/repository"
	"final-project-pemin/util"
)

type (
	IMahasiswaService interface {
		Create(req *model.MahasiswaInputRequest) (*model.Mahasiswa, error)
		Login(req *model.MahasiswaLoginRequest) (*model.Mahasiswa, error)
		FindByNIM(nim string) (*model.Mahasiswa, error)
		FindAll() ([]*model.Mahasiswa, error)
	}

	mahasiswaService struct {
		mahasiswaRepository repository.IMahasiswaRepository
	}
)

func NewMahasiswaService(mahasiswaRepository repository.IMahasiswaRepository) IMahasiswaService {
	return &mahasiswaService{mahasiswaRepository: mahasiswaRepository}
}

func (ms *mahasiswaService) Create(req *model.MahasiswaInputRequest) (*model.Mahasiswa, error) {
	err := util.ValidationPassword(req.Password)
	if err != nil {
		return nil, err
	}

	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	mahasiswa := &model.Mahasiswa{
		NIM:      req.NIM,
		ProdiId:  req.ProdiId,
		Nama:     req.Nama,
		Angkatan: req.Angkatan,
		Password: hashPassword,
	}

	newMahasiswa, err := ms.mahasiswaRepository.Save(mahasiswa)
	if err != nil {
		return nil, err
	}

	return newMahasiswa, nil
}

func (ms *mahasiswaService) Login(req *model.MahasiswaLoginRequest) (*model.Mahasiswa, error) {
	mahasiswa, err := ms.mahasiswaRepository.GetByNIM(req.NIM)
	if err != nil {
		return nil, errors.New("invalid nim or password")
	}

	err = util.ComparePassword(mahasiswa.Password, req.Password)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return mahasiswa, nil
}

func (ms *mahasiswaService) FindByNIM(nim string) (*model.Mahasiswa, error) {
	mahasiswa, err := ms.mahasiswaRepository.GetByNIM(nim)
	if err != nil {
		return nil, err
	}

	return mahasiswa, nil
}

func (ms *mahasiswaService) FindAll() ([]*model.Mahasiswa, error) {
	mahasiswas, err := ms.mahasiswaRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return mahasiswas, nil
}
