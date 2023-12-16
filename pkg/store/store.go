package store

import (
	"errors"
	"web-endterm/internal/model"
)

type Store interface {
	CreatePost(post *model.Post) error
	GetPostById(id int) (*model.Post, error)
	GetAllPosts()([]*model.Post, error)
	PatchPost(id int, post *model.Post) (*model.Post, error)
	DeletePost(id int) error
}

type store struct {
	posts []*model.Post
}

var Posts []*model.Post

func NewStore(posts []*model.Post) Store{
	return &store{
		posts: Posts,
	}
}

func (s *store) CreatePost(post *model.Post) error {
	s.posts = append(s.posts, post)

	return nil
}

func (s *store) GetPostById(id int) (*model.Post, error) {
	for i, _ := range s.posts{
		if i == id {
			return s.posts[i], nil
		}
	}

	return nil, errors.New("no post with such id")
}

func (s *store) GetAllPosts()([]*model.Post, error) {
	return s.posts, nil
}

func (s *store) PatchPost(id int, post *model.Post) (*model.Post, error) {
	for i, _ := range s.posts{
		if i == id {
			if post.Author != ""{
				s.posts[i].Author = post.Author
			}
			if post.Title != ""{
				s.posts[i].Title = post.Title
			}
			if post.Text != ""{
				s.posts[i].Text = post.Text
			}
			if post.DateOfCreation != ""{
				s.posts[i].DateOfCreation = post.DateOfCreation
			}

			return s.posts[i], nil
		}
	}

	return nil, errors.New("no post with such id")
}

func (s *store) DeletePost(id int) error {
	for i, _ := range s.posts{
		if i == id {
			s.posts = RemoveIndex(s.posts, id)
			return nil
		}
	}

	return errors.New("no post  with such id")
}

func RemoveIndex(s []*model.Post, index int) []*model.Post {
    ret := make([]*model.Post, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}