package service

import (
	"errors"
	"final-project-pemin/src/model"
	"final-project-pemin/src/repository"
	"final-project-pemin/util"
	"fmt"
	"strconv"
)

type (
	IMahasiswaService interface {
		Create(req *model.MahasiswaInputRequest) (*model.Mahasiswa, error)
		Login(req *model.MahasiswaLoginRequest) (*model.Mahasiswa, error)
		FindByNIM(nim string) (*model.Mahasiswa, error)
		FindAll() ([]*model.Mahasiswa, error)
		SaveMatkul(nim string, mkId int64) (*model.Mahasiswa, error)
		DeleteMatkul(nim string, mkId int64) (*model.Mahasiswa, error)
	}

	mahasiswaService struct {
		mahasiswaRepository repository.IMahasiswaRepository
		matkulRepository    repository.IMataKuliahRepository
	}
)

func NewMahasiswaService(mahasiswaRepository repository.IMahasiswaRepository, matkulRepository repository.IMataKuliahRepository) IMahasiswaService {
	return &mahasiswaService{
		mahasiswaRepository: mahasiswaRepository,
		matkulRepository:    matkulRepository,
	}
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

	prodiId, _ := strconv.Atoi(req.ProdiId)
	angkatan, _ := strconv.Atoi(req.Angkatan)

	mahasiswa := &model.Mahasiswa{
		NIM:      req.NIM,
		ProdiId:  int64(prodiId),
		Nama:     req.Nama,
		Angkatan: int32(angkatan),
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

func (ms *mahasiswaService) SaveMatkul(nim string, mkId int64) (*model.Mahasiswa, error) {
	mahasiswa, err := ms.mahasiswaRepository.GetByNIM(nim)
	if err != nil {
		return nil, err
	}

	matkul, err := ms.matkulRepository.GetByID(mkId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var exist bool
	for _, matkulMhs := range mahasiswa.MataKuliah {
		if matkulMhs.ID == matkul.ID {
			exist = true
		}
	}

	if exist {
		return nil, errors.New("Mata kuliah " + matkul.Nama + " sudah ditambahkan")
	}

	err = ms.mahasiswaRepository.SaveMatkul(mahasiswa, matkul)
	if err != nil {
		return nil, err
	}

	return mahasiswa, nil
}

func (ms *mahasiswaService) DeleteMatkul(nim string, mkId int64) (*model.Mahasiswa, error) {
	mahasiswa, err := ms.mahasiswaRepository.GetByNIM(nim)
	if err != nil {
		return nil, err
	}

	matkul, err := ms.matkulRepository.GetByID(mkId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = ms.mahasiswaRepository.DeleteMatkul(mahasiswa, matkul)
	if err != nil {
		return nil, err
	}

	return mahasiswa, nil
}
