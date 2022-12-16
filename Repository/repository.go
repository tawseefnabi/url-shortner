package repository

import (
	"log"
	model "url-shortner/Model"

	"github.com/google/uuid"
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

func (r *Repository) Save(urlModel model.UrlModel, hash string) {
	var urlData model.TinyUrlData
	r.Db.Where("hash = (?)", hash).Find(&urlData)
	if urlData.Url == "" {
		r.Db.Create(&model.TinyUrlData{
			Hash: hash,
			Url:  urlModel.Url,
		})
		log.Println("Data is created for url: ", urlModel.Url, " with hash: ", hash)
	} else {
		id := uuid.New()
		r.Db.Create(&model.TinyUrlData{
			Hash: id.String(),
			Url:  urlModel.Url,
		})
		log.Println("hash already existed, creating another hash")
	}
}

func (r *Repository) Get(hash string) model.TinyUrlData {
	var urlData model.TinyUrlData
	r.Db.Where("hash = (?)", hash).Find(&urlData)
	return urlData
}
