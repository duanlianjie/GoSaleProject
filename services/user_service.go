package services

import (
	"golang.org/x/crypto/bcrypt"
	"goproject/datamodels"
	"goproject/repositories"
)

type UserService interface {
	IsPwdSuccess(userName string, pwd string) (user *datamodels.User, isOK bool)
	AddUser(user *datamodels.User) (userID int64, err error)
}

type UserServiceManager struct {
	UserRepository repositories.UserRepository
}

func (u *UserServiceManager) IsPwdSuccess(userName string, pwd string) (user *datamodels.User, isOK bool) {
	//panic("implement me")
	user, err := u.UserRepository.Select(userName)
	if err != nil {
		return &datamodels.User{}, false
	}
	isOK, _ = ValidatePassword(pwd, user.HashPassword)
	if !isOK {
		return &datamodels.User{}, false
	}
	return
}

func (u *UserServiceManager) AddUser(user *datamodels.User) (userID int64, err error) {
	//panic("implement me")
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.Insert(user)
}

func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashPassword string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(userPassword)); err != nil {
		return false, err
	}
	return true, nil
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &UserServiceManager{UserRepository: repository}
}
