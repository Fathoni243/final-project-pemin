package repository

import (
	"final-project-pemin/src/model"

	"gorm.io/gorm"
)

type (
	IProdiRepository interface {
		GetAll() ([]*model.Prodi, error)
	}

	prodiRepository struct {
		db *gorm.DB
	}
)

func NewProdiRepository(db *gorm.DB) IProdiRepository {
	return &prodiRepository{db: db}
}

func (pr *prodiRepository) GetAll() ([]*model.Prodi, error) {
	tx := pr.db.Begin()

	var prodis []*model.Prodi
	err := tx.Find(&prodis).Error
	if err != nil {
		return nil, err
	}

	return prodis, nil
}

