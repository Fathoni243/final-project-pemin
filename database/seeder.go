package database

import "final-project-pemin/src/model"

func SeederRefresh() {
	db, _ := InitMySQL()

	db.Migrator().DropTable(
		&model.Prodi{},
		&model.Mahasiswa{},
		&model.MataKuliah{},
	)

	db.AutoMigrate(
		&model.Prodi{},
		&model.Mahasiswa{},
		&model.MataKuliah{},
	)

	db.Model(&model.Prodi{}).Create([]map[string]interface{}{
		{"Nama": "Teknologi Informasi"},
		{"Nama": "Sistem Informasi"},
		{"Nama": "Pendidikan Teknologi Informasi"},
		{"Nama": "Teknik Informatika"},
		{"Nama": "Teknik Komputer"},
	})

	db.Model(&model.MataKuliah{}).Create([]map[string]interface{}{
		{"Nama": "Pemrograman Dasar"},
		{"Nama": "Pemrograman Lanjut"},
		{"Nama": "Algoritma dan Struktur Data"},
		{"Nama": "Sistem Basis Data"},
		{"Nama": "Jaringan Komputer Dasar"},
	})
}
