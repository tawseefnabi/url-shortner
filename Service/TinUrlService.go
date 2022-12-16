package service

import (
	model "url-shortner/Model"
	repository "url-shortner/Repository"
	utility "url-shortner/Utility"
)

var (
	UrlContent = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	UrlAddress = "localhost:8081/"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
func (s *Service) GenerateTinyUrl(urlModel model.UrlModel) model.UrlModel {
	hash := utility.ComputeHash(urlModel.Url)
	len := len(UrlContent)
	hashUrl := ""
	for hash > 0 {
		idx := hash % int64(len)
		hash = hash / int64(len)
		hashUrl += string(UrlContent[idx])
	}
	computedUrl := UrlAddress + hashUrl
	s.repo.Save(urlModel, hashUrl)

	// computedUrl = "hello"
	return model.UrlModel{
		Url: computedUrl,
	}
}

func (s *Service) RedirectTinyUrl(hash string) string {
	tinyUrl := s.repo.Get(hash)
	return tinyUrl.Url
}
