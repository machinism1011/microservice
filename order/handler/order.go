package handler

import (
	"context"

	"github.com/machinism1011/microservice/order/domain/model"

	"github.com/machinism1011/microservice/common"
	"github.com/machinism1011/microservice/order/domain/service"
	protoOrder "github.com/machinism1011/microservice/order/proto/order"
)

type Order struct {
	OrderDataService service.IOrderDataService
}

func (o *Order) GetOrderByID(_ context.Context, request *protoOrder.OrderID, response *protoOrder.OrderInfo) error {
	order, err := o.OrderDataService.FindOrderByID(request.OrderId)
	if err != nil {
		return err
	}
	if err := common.SwapTo(order, response); err != nil {
		return err
	}
	return nil
}

func (o *Order) GetAllOrder(_ context.Context, _ *protoOrder.AllOrderRequest, response *protoOrder.AllOrder) error {
	orderSlice, err := o.OrderDataService.FindAllOrder()
	if err != nil {
		return err
	}
	orderItem := &protoOrder.OrderInfo{}
	for _, order := range orderSlice {
		if err := common.SwapTo(order, orderItem); err != nil {
			return err
		}
		response.OrderInfo = append(response.OrderInfo, orderItem)
	}
	return nil
}

func (o *Order) CreateOrder(_ context.Context, request *protoOrder.OrderInfo, response *protoOrder.OrderID) error {
	order := &model.Order{}
	if err := common.SwapTo(request, order); err != nil {
		return err
	}
	orderID, err := o.OrderDataService.AddOrder(order)
	if err != nil {
		return err
	}
	response.OrderId = orderID
	return nil
}

func (o *Order) DeleteOrderByID(_ context.Context, request *protoOrder.OrderID, response *protoOrder.Response) error {
	err := o.OrderDataService.DeleteOrder(request.OrderId)
	if err != nil {
		response.StatusCode = 500
		response.Message = "删除失败"
	} else {
		response.StatusCode = 200
		response.Message = "删除成功"
	}
	return err
}

func (o *Order) UpdateOrderPayStatus(_ context.Context, request *protoOrder.PayStatus, response *protoOrder.Response) error {
	if err := o.OrderDataService.UpdatePayStatus(request.OrderId, request.PayStatus); err != nil {
		response.StatusCode = 500
		response.Message = "更新支付状态失败"
		return err
	} else {
		response.StatusCode = 200
		response.Message = "更新支付状态成功"
		return nil
	}
}

func (o *Order) UpdateOrderShipStatus(_ context.Context, request *protoOrder.ShipStatus, response *protoOrder.Response) error {
	if err := o.OrderDataService.UpdatePayStatus(request.OrderId, request.ShipStatus); err != nil {
		response.StatusCode = 500
		response.Message = "更新发货状态失败"
		return err
	} else {
		response.StatusCode = 200
		response.Message = "更新发货状态成功"
		return nil
	}
}

func (o *Order) UpdateOrder(_ context.Context, request *protoOrder.OrderInfo, response *protoOrder.Response) error {
	order := &model.Order{}
	if err := common.SwapTo(request, order); err != nil {
		return err
	}
	if err := o.OrderDataService.UpdateOrder(order); err != nil {
		response.StatusCode = 500
		response.Message = "更新订单失败"
		return err
	} else {
		response.StatusCode = 200
		response.Message = "更新订单成功"
		return nil
	}
}
