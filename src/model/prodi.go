package model

type (
	Prodi struct {
		ID        int64       `json:"id" gorm:"primaryKey"`
		Nama      string      `json:"nama" gorm:"not null;type:varchar(100)"`
		Mahasiswa []Mahasiswa `json:"-" gorm:"foreignKey:ProdiId"`
	}
)
