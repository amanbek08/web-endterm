package service

import (
	"web-endterm/internal/model"
	"web-endterm/pkg/store"
)

type IService interface {
	CreatePost(post *model.Post) error
	GetPostById(id int) (*model.Post, error)
	GetAllPosts()([]*model.Post, error)
	PatchPost(id int, post *model.Post) (*model.Post, error)
}

type Service struct {
	store store.Store
}

func NewService(s store.Store) Service{
	return Service{
		store: s,
	}
}

func (s *Service) CreatePost(post *model.Post)error{
	err := s.store.CreatePost(post)
	
	return err
}

func (s *Service) GetPostById(id int) (*model.Post, error){
	post, err := s.store.GetPostById(id)

	return post, err
}

func (s *Service) GetAllPosts()([]*model.Post, error) {
	posts, err := s.store.GetAllPosts()

	return posts, err
}

func (s *Service) PatchPost(id int, post *model.Post) (*model.Post, error) {
	post, err := s.store.PatchPost(id, post)

	return post, err
}

func (s *Service) DeletePost(id int) error {
	return s.store.DeletePost(id)
}