package repository

import (
	model "url-shortner/Model"

	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	// Migrate the schema
	db.AutoMigrate(&model.TinyUrlData{})
	return &Repository{
		Db: db,
	}
}
