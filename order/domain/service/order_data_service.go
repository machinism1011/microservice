package service

import (
	"github.com/machinism1011/microservice/order/domain/model"
	"github.com/machinism1011/microservice/order/domain/repository"
)

type IOrderDataService interface {
	AddOrder(*model.Order) (int64, error)
	DeleteOrder(int64) error
	UpdateOrder(*model.Order) error
	FindOrderByID(int64) (*model.Order, error)
	FindAllOrder() ([]model.Order, error)
	UpdateShipStatus(int64, int32) error
	UpdatePayStatus(int64, int32) error
}

type OrderDataService struct {
	OrderRepository repository.IOrderRepository
}

func NewOrderDataService(orderRepository repository.IOrderRepository) IOrderDataService {
	return &OrderDataService{orderRepository}
}

func (o *OrderDataService) AddOrder(order *model.Order) (int64, error) {
	return o.OrderRepository.CreateOrder(order)
}

func (o *OrderDataService) DeleteOrder(orderID int64) error {
	return o.OrderRepository.DeleteOrderByID(orderID)
}

func (o *OrderDataService) UpdateOrder(order *model.Order) error {
	return o.OrderRepository.UpdateOrder(order)
}

func (o *OrderDataService) FindOrderByID(orderID int64) (*model.Order, error) {
	return o.OrderRepository.FindOrderByID(orderID)
}

func (o *OrderDataService) FindAllOrder() ([]model.Order, error) {
	return o.OrderRepository.FindAllOrder()
}

func (o *OrderDataService) UpdateShipStatus(orderID int64, shipStatus int32) error {
	return o.OrderRepository.UpdateShipStatus(orderID, shipStatus)
}

func (o *OrderDataService) UpdatePayStatus(orderID int64, payStatus int32) error {
	return o.OrderRepository.UpdatePayStatus(orderID, payStatus)
}
