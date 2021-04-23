package services

import (
	"fmt"
	"log"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/repository"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/request"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthServices interface {
	CreateUser(req request.RequestAuthRegister) (entity.Users, error)
	UserIsExist(req string) bool
	VerifyCredential(rusername string, password string) interface{}
	EmailIsExist(req request.RequestAuthForgetPassword) bool
}

type authServices struct {
	repository repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthServices {
	return &authServices{
		repository: repo,
	}
}

func (s *authServices) CreateUser(req request.RequestAuthRegister) (entity.Users, error) {
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
	fmt.Println(hashedPassword)
	newUser, err := s.repository.InsertUser(user)
	if err != nil {
		return user, err
	}
	return newUser, nil
}

func (s *authServices) UserIsExist(req string) bool {
	result := s.repository.UserIsExist(req)
	return result
}

func (s *authServices) VerifyCredential(username string, password string) interface{} {
	res := s.repository.GetByUsername(username)
	if v, ok := res.(entity.Users); ok {
		comparedPassword := comparePassword(v.Password, password)
		fmt.Println("password :", password)
		fmt.Println("Password :", v.Password)
		fmt.Println("Username :", v.Username)
		fmt.Println("username :", username)
		fmt.Println("ComparedPassword :", comparedPassword)
		fmt.Println("byte password :", []byte(password))
		fmt.Println("byte Password :", []byte(v.Password))
		if v.Username == username && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (s *authServices) UpdatePassword(req request.RequestAuthUpdate) bool {
	var user entity.Users
	generatePassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	user.ID = req.ID
	user.Password = string(generatePassword)
	res := s.repository.UpdateUserPassword(user)
	return res == nil
}

func (s *authServices) EmailIsExist(req request.RequestAuthForgetPassword) bool {
	res := s.repository.EmailIsExist(req.Email)
	return res
}

func comparePassword(hashedPwd string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPassword))
	return err == nil
}
