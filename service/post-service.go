package service

import (
	"errors"
	"github.com/ahmedkhaeld/rest-api/entity"
	"github.com/ahmedkhaeld/rest-api/repository"
	"math/rand"
)

var (
	repo repository.PostRepository
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

// create a service struct, that implements the service interface
type service struct {
}

// NewPostService constructor
func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("the post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("the post title is empty")
		return err
	}

	return nil

}
func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)

}
func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
