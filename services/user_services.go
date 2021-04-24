package services

import (
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/repository"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/request"
)

type UserServices interface {
	Update(req request.RequestUser) (entity.Users, error)
	Profile(req request.RequestUserProfile) (entity.Users, error)
}

type userServices struct {
	repository repository.UserRepository
}

func NewUserServices(repository repository.UserRepository) *userServices {
	return &userServices{repository}
}

func (s *userServices) Update(req request.RequestUser) (entity.Users, error) {
	user := entity.Users{}
	user.ID = req.ID
	user.Username = req.Username
	user.Fullname = req.Fullname
	user.Email = req.Email
	user.Role = req.Role
	updated, err := s.repository.UpdateUser(user)
	if err != nil {
		return user, err
	}
	return updated, nil
}

func (s *userServices) Profile(req request.RequestUserProfile) (entity.Users, error) {
	user := entity.Users{}
	user.ID = req.ID
	prof, err := s.repository.ProfileUser(user)
	if err != nil {
		return user, err
	}
	return prof, nil
}
