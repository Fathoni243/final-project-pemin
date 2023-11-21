package model

type (
	MataKuliah struct {
		ID   int64  `json:"id" gorm:"primaryKey"`
		Nama string `json:"nama" gorm:"not null;type:varchar(100)"`
	}
)
