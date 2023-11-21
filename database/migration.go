package database

import (
	"final-project-pemin/src/model"

	"gorm.io/gorm"
)

type Migration struct {
	DB *gorm.DB
}

func (m *Migration) RunMigration() {
	m.DB.Migrator().DropTable(
		&model.Prodi{},
		&model.Mahasiswa{},
		&model.MataKuliah{},
	)

	m.DB.AutoMigrate(
		&model.Prodi{},
		&model.Mahasiswa{},
		&model.MataKuliah{},
	)
}
