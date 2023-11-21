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
	err := tx.Preload("Prodi").First(mahasiswa, nim).Error
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
