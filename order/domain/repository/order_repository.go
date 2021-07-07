package repository

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/machinism1011/microservice/order/domain/model"
)

type IOrderRepository interface {
	InitTable() error
	FindOrderByID(int64) (*model.Order, error)
	FindAllOrder() ([]model.Order, error)
	CreateOrder(*model.Order) (int64, error)
	DeleteOrderByID(int64) error
	UpdateOrder(*model.Order) error
	UpdateShipStatus(int64, int32) error
	UpdatePayStatus(int64, int32) error
}

type OrderRepository struct {
	mysqlDb *gorm.DB
}

func NewOrderRepository(db *gorm.DB) IOrderRepository {
	return &OrderRepository{mysqlDb: db}
}

func (o *OrderRepository) InitTable() error {
	// create multi tables
	return o.mysqlDb.CreateTable(&model.Order{}, &model.OrderDetail{}).Error
}

func (o *OrderRepository) FindOrderByID(orderID int64) (order *model.Order, err error) {
	order = &model.Order{}
	return order, o.mysqlDb.Preload("OrderDetail").First(order, orderID).Error
}

func (o *OrderRepository) FindAllOrder() (orderSlice []model.Order, err error) {
	orderSlice = make([]model.Order, 0)
	return orderSlice, o.mysqlDb.Preload("OrderDetail").Find(&orderSlice).Error
}

func (o *OrderRepository) CreateOrder(order *model.Order) (orderID int64, err error) {
	return order.ID, o.mysqlDb.Create(order).Error
}

func (o *OrderRepository) DeleteOrderByID(orderID int64) error {
	// 两个表 开启事物
	tx := o.mysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Unscoped().Where("id = ?", orderID).Delete(&model.Order{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Where("order_id = ?", orderID).Delete(&model.Order{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (o *OrderRepository) UpdateOrder(order *model.Order) error {
	return o.mysqlDb.Model(order).Update(order).Error
}

func (o *OrderRepository) UpdateShipStatus(orderID int64, shipStatus int32) error {
	db := o.mysqlDb.Model(&model.Order{}).Where("id = ?", orderID).UpdateColumn("ship_status", shipStatus)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return nil
}

func (o *OrderRepository) UpdatePayStatus(orderID int64, payStatus int32) error {
	db := o.mysqlDb.Model(&model.Order{}).Where("id = ?", orderID).UpdateColumn("pay_status", payStatus)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return nil
}
