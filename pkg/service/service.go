package service

import "github.com/higatu/todo-app/pkg/repository"


type Banner interface{

}

type Tag interface{

}

type Feature interface{

}

type Service struct {
	Banner
	Tag
	Feature
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
