package repository

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/machinism1011/microservice/cart/domain/model"
)

type ICartRepository interface {
	InitTable() error
	FindCartByID(int64) (*model.Cart, error)
	FindAllCart(int64) ([]model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error

	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

type CartRepository struct {
	mysqlDb *gorm.DB
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDb: db}
}

func (c *CartRepository) InitTable() error {
	// create multi tables
	return c.mysqlDb.CreateTable(&model.Cart{}).Error
}

func (c *CartRepository) FindCartByID(cartID int64) (cart *model.Cart, err error) {
	cart = &model.Cart{}
	return cart, c.mysqlDb.First(cart, cartID).Error
}

func (c *CartRepository) FindAllCart(userID int64) (cartSlice []model.Cart, err error) {
	cartSlice = make([]model.Cart, 0)
	return cartSlice, c.mysqlDb.Where("user_id = ?", userID).Find(&cartSlice).Error
}

func (c *CartRepository) CreateCart(cart *model.Cart) (cartID int64, err error) {
	// 根据ProductID和SizeID判断是否存在，如果存在则不创建，否则创建。
	db := c.mysqlDb.FirstOrCreate(cart, model.Cart{ProductID: cart.ProductID, SizeID: cart.SizeID, UserID: cart.UserID})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("购物车插入失败")
	}
	return cart.ID, nil
}

func (c *CartRepository) DeleteCartByID(cartID int64) error {
	return c.mysqlDb.Where("id = ?", cartID).Delete(&model.Cart{}).Error
}

func (c *CartRepository) UpdateCart(cart *model.Cart) error {
	return c.mysqlDb.Model(cart).Update(cart).Error
}

func (c *CartRepository) CleanCart(userID int64) error {
	return c.mysqlDb.Where("user_id = ?", userID).Delete(&model.Cart{}).Error
}

func (c *CartRepository) IncrNum(cartID, num int64) error {
	cart := &model.Cart{ID: cartID}
	return c.mysqlDb.Model(cart).UpdateColumn("num", gorm.Expr("num + ?", num)).Error
}

func (c *CartRepository) DecrNum(cartID, num int64) error {
	cart := &model.Cart{ID: cartID}
	db := c.mysqlDb.Model(cart).Where("num >= ?", num).UpdateColumn("num", gorm.Expr("num - ?", num))
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("减少失败")
	}
	return nil

}
