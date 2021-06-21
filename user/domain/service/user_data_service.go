package service

import (
	"errors"
	"github.com/machinism1011/microservice/user/domain/model"
	"github.com/machinism1011/microservice/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

// 在repository的基础上实现业务逻辑代码
type IUserDataService interface {
	AddUser(*model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(*model.User, bool) error
	FindUserByName(string) (*model.User, error)
	CheckPwd(userName string, pwd string) (isOk bool, err error)
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository:userRepository}
}
// 加密及验证用户密码
func generatePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func validatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误")
	}
	return true, nil
}

// 插入用户
func (u *UserDataService) AddUser(user *model.User) (userID int64, err error) {
	pwdByte, err := generatePassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}
	// 存储加密后的用户密码
	user.HashPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}

func (u *UserDataService) DeleteUser(userID int64) error {
	return u.UserRepository.DeleteUserByID(userID)
}

func (u *UserDataService) UpdateUser(user *model.User, isChangedPwd bool) (err error) {
	if isChangedPwd {
		pwdByte, err := generatePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	return u.UserRepository.UpdateUser(user)
}

func (u *UserDataService) FindUserByName(userName string) (user *model.User, err error) {
	return u.UserRepository.FindUserByName(userName)
}

func (u *UserDataService) CheckPwd(userName, pwd string) (isOk bool, err error) {
	user, err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false, err
	}
	return validatePassword(pwd, user.HashPassword)
}

