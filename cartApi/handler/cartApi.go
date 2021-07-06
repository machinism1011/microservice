package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/micro/go-micro/v2/logger"

	protoCart "github.com/machinism1011/microservice/cart/proto/cart"
	protoCartApi "github.com/machinism1011/microservice/cartApi/proto/cartApi"
)

type CartApi struct {
	CartService protoCart.CartService
}

func (c *CartApi) FindAll(_ context.Context, request *protoCartApi.Request, response *protoCartApi.Response) error {
	logger.Info("接收到 /cartApi/findAll 访问请求")
	if _, ok := request.Get["user_id"]; !ok {
		return errors.New("参数异常")
	}
	userIDString := request.Get["user_id"].Values[0]
	fmt.Println(userIDString)
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		return err
	}

	// 获取购物车所有商品
	cartAll, err := c.CartService.GetAll(context.TODO(), &protoCart.CartFindAll{UserId: userID})
	b, err := json.Marshal(cartAll)
	if err != nil {
		return err
	}
	response.StatusCode = 200
	response.Body = string(b)
	return nil
}
