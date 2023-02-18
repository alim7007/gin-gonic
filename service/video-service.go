package service

import "github.com/alim7007/gin-gonic/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type serviceVideo struct {
	videos []entity.Video
}

func New() VideoService{
	return &serviceVideo{
		videos : []entity.Video{},
	}
}

func(service *serviceVideo) Save(video entity.Video) entity.Video{
	service.videos = append(service.videos, video)
	return video
}

func(service *serviceVideo) FindAll() []entity.Video{
	return service.videos
}