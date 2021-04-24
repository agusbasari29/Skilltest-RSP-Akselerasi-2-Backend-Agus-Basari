package services

import (
	"fmt"
	"log"
	"time"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/repository"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/request"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	CreateUser(req request.RequestAuthRegister) (entity.Users, error)
	UserIsExist(req string) bool
	VerifyCredential(rusername string, password string) interface{}
	EmailIsExist(req request.RequestAuthForgetPassword) bool
	Update(req request.RequestUser) (entity.Users, error)
	UpdatePassword(req request.RequestAuthUpdate) bool
	Profile(req request.RequestUserProfile) (entity.Users, error)
	GetAllCreator(creator entity.Users) ([]entity.Users, error)
	DeleteCreator(creator entity.Users) error
}

type userServices struct {
	repository repository.UserRepository
}

func NewUserServices(repository repository.UserRepository) *userServices {
	return &userServices{repository}
}

func (s *userServices) Update(req request.RequestUser) (entity.Users, error) {
	user := entity.Users{}
	err := smapping.FillStruct(&user, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}
	user.CreatedAt = time.Now()
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

func (s *userServices) CreateUser(req request.RequestAuthRegister) (entity.Users, error) {
	user := entity.Users{}
	err := smapping.FillStruct(&user, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	fmt.Println(hashedPassword)
	newUser, err := s.repository.InsertUser(user)
	if err != nil {
		return user, err
	}
	return newUser, nil
}

func (s *userServices) UserIsExist(req string) bool {
	result := s.repository.UserIsExist(req)
	return result
}

func (s *userServices) VerifyCredential(username string, password string) interface{} {
	res := s.repository.GetByUsername(username)
	if v, ok := res.(entity.Users); ok {
		comparedPassword := comparePassword(v.Password, password)
		if v.Username == username && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (s *userServices) UpdatePassword(req request.RequestAuthUpdate) bool {
	var user entity.Users
	generatePassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	user.ID = req.ID
	user.Password = string(generatePassword)
	user.UpdatedAt = time.Now()
	res := s.repository.UpdateUserPassword(user)
	return res == nil
}

func (s *userServices) EmailIsExist(req request.RequestAuthForgetPassword) bool {
	res := s.repository.EmailIsExist(req.Email)
	return res
}

func (s *userServices) GetAllCreator(creator entity.Users) ([]entity.Users, error) {
	res, err := s.repository.GetUserByRole(creator)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *userServices) DeleteCreator(creator entity.Users) error {
	err := s.repository.DeleteCreator(creator)
	if err != nil {
		return err
	}
	return nil
}

func comparePassword(hashedPwd string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPassword))
	return err == nil
}
