package repository

import (
	"final-project-pemin/src/model"
	"final-project-pemin/util"

	"gorm.io/gorm"
)

type (
	IMataKuliahRepository interface {
		GetAll() ([]*model.MataKuliah, error)
	}

	mataKuliahRepository struct {
		db *gorm.DB
	}
)

func NewMataKuliahRepository(db *gorm.DB) IMataKuliahRepository {
	return &mataKuliahRepository{db: db}
}

func (mkr *mataKuliahRepository) GetAll() ([]*model.MataKuliah, error) {
	tx := mkr.db.Begin()

	var matakuliahs []*model.MataKuliah
	err := tx.Find(&matakuliahs).Error
	if err != nil {
		return nil, err
	}

	err = util.CommitOrRollback(tx)
	if err != nil {
		return nil, err
	}

	return matakuliahs, nil

}