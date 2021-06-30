package handler

import (
	"context"

	"github.com/machinism1011/microservice/cart/domain/model"
	"github.com/machinism1011/microservice/cart/domain/service"
	proto "github.com/machinism1011/microservice/cart/proto/cart"
	"github.com/machinism1011/microservice/common"
)

type Cart struct {
	CartDataService service.ICartDataService
}

func (c *Cart) AddCart(_ context.Context, request *proto.CartInfo, response *proto.ResponseAdd) (err error) {
	cart := &model.Cart{}
	if err = common.SwapTo(request, cart); err != nil {
		return err
	}
	response.CartId, err = c.CartDataService.AddCart(cart)
	return err
}

func (c *Cart) ClearCart(_ context.Context, request *proto.Clean, response *proto.Response) error {
	if err := c.CartDataService.CleanCart(request.UserId); err != nil {
		return err
	}
	response.Message = "购物车清空成功"
	return nil
}

func (c *Cart) Incr(_ context.Context, request *proto.Item, response *proto.Response) error {
	if err := c.CartDataService.IncrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Message = "添加到购物车成功"
	return nil
}

func (c *Cart) Decr(_ context.Context, request *proto.Item, response *proto.Response) error {
	if err := c.CartDataService.DecrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Message = "购物车减少数量成功"
	return nil
}

func (c *Cart) DeleteItemByID(_ context.Context, request *proto.CartID, response *proto.Response) error {
	if err := c.CartDataService.DeleteCart(request.Id); err != nil {
		return err
	}
	response.Message = "购物车项目删除成功"
	return nil
}
func (c *Cart) GetAll(_ context.Context, request *proto.CartFindAll, response *proto.CartAll) error {
	cartSlice, err := c.CartDataService.FindAllCart(request.UserId)
	if err != nil {
		return err
	}
	cart := &proto.CartInfo{}
	for _, v := range cartSlice {
		if err := common.SwapTo(v, cart); err != nil {
			return err
		}
		response.CartInfo = append(response.CartInfo, cart)
	}
	return nil
}
