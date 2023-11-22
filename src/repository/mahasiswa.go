package repository

import (
	"final-project-pemin/src/model"
	"final-project-pemin/util"

	"gorm.io/gorm"
)

type (
	IMahasiswaRepository interface {
		Save(mahasiswa *model.Mahasiswa) (*model.Mahasiswa, error)
		GetByNIM(nim string) (*model.Mahasiswa, error)
		GetAll() ([]*model.Mahasiswa, error)
		SaveMatkul(mahasiswa *model.Mahasiswa, matkul *model.MataKuliah) error
		DeleteMatkul(mahasiswa *model.Mahasiswa, matkul *model.MataKuliah) error
	}

	mahasiswaRepository struct {
		db *gorm.DB
	}
)

func NewMahasiswaRepository(db *gorm.DB) IMahasiswaRepository {
	return &mahasiswaRepository{
		db: db,
	}
}

func (mr *mahasiswaRepository) Save(mahasiswa *model.Mahasiswa) (*model.Mahasiswa, error) {
	tx := mr.db.Begin()

	err := tx.Create(&mahasiswa).Error
	if err != nil {
		return nil, err
	}

	err = util.CommitOrRollback(tx)
	if err != nil {
		return nil, err
	}

	return mahasiswa, nil
}

func (mr *mahasiswaRepository) GetByNIM(nim string) (*model.Mahasiswa, error) {
	tx := mr.db.Begin()

	mahasiswa := new(model.Mahasiswa)
	err := tx.Preload("Prodi").Preload("MataKuliah").First(mahasiswa, nim).Error
	if err != nil {
		return nil, err
	}

	err = util.CommitOrRollback(tx)
	if err != nil {
		return nil, err
	}

	return mahasiswa, nil
}

func (mr *mahasiswaRepository) GetAll() ([]*model.Mahasiswa, error) {
	tx := mr.db.Begin()

	var mahasiswas []*model.Mahasiswa
	err := tx.Preload("Prodi").Find(&mahasiswas).Error
	if err != nil {
		return nil, err
	}

	err = util.CommitOrRollback(tx)
	if err != nil {
		return nil, err
	}

	return mahasiswas, nil
}

func (mr *mahasiswaRepository) SaveMatkul(mahasiswa *model.Mahasiswa, matkul *model.MataKuliah) error {
	tx := mr.db.Begin()

	err := tx.Model(&mahasiswa).Preload("MataKuliah").Association("MataKuliah").Append(matkul)
	if err != nil {
		return err
	}

	err = util.CommitOrRollback(tx)
	if err != nil {
		return err
	}

	return nil
}

func (mr *mahasiswaRepository) DeleteMatkul(mahasiswa *model.Mahasiswa, matkul *model.MataKuliah) error {
	tx := mr.db.Begin()

	err := tx.Model(&mahasiswa).Preload("MataKuliah").Association("MataKuliah").Delete(matkul)
	if err != nil {
		return err
	}

	err = util.CommitOrRollback(tx)
	if err != nil {
		return err
	}

	return nil
}
