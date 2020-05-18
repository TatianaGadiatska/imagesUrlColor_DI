package service

import (
	st "../data_struct"
	db "../database"
	f "../find_images"
)

type UrlColorService struct {
	config     *st.Config
	repository *db.UrlRepository
	urlAll     *f.Result
}

func (service *UrlColorService) FindAll() []st.UrlImage {
	if service.config.Enabled {
		return service.repository.FindAll()
	}
	return []st.UrlImage{}
}

func (service *UrlColorService) CreateTable() {
	service.repository.CreateTable()
}

func (service *UrlColorService) InsertUrl(urlColorBd []st.UrlImage) {
	service.repository.InsertUrl(urlColorBd)
}

func (service *UrlColorService) MakeUrlColor() []st.UrlImage {
	return service.urlAll.MakeUrlColor()
}

func NewUrlColorService(config *st.Config, repository *db.UrlRepository) *UrlColorService {
	return &UrlColorService{config: config, repository: repository}
}
