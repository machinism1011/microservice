package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/machinism1011/microservice/user/domain/model"
)

// user data与数据库交互的接口
type IUserRepository interface {
	// 初始化数据表
	InitTable() error
	// 根据用户名及ID查找用户信息
	FindUserByName(string) (*model.User, error)
	FindUserByID(int64) (*model.User, error)
	// 增删改
	CreateUser(user *model.User) (int64, error)
	DeleteUserByID(int64) error
	UpdateUser(*model.User) error
	// 查找所有用户
	FindAll() ([]model.User, error)
}

// 实现IUserRepository接口
type UserRepository struct {
	mysqlDb *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		mysqlDb: db,
	}
}

func (u *UserRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.User{}).Error
}

func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.Where("user_name=?", name).Find(user).Error
}

func (u *UserRepository) FindUserByID(userID int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.First(user, userID).Error
}

func (u *UserRepository) CreateUser(user *model.User) (int64, error) {
	return user.ID, u.mysqlDb.Create(user).Error
}

func (u *UserRepository) DeleteUserByID(userID int64) error {
	return u.mysqlDb.Where("id = ?", userID).Delete(&model.User{}).Error
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Update(&user).Error
}

func (u *UserRepository) FindAll() (userAll []model.User, err error) {
	return userAll, u.mysqlDb.Find(&userAll).Error
}