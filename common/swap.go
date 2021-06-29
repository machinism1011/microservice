package common

import (
	"encoding/json"

	"github.com/machinism1011/microservice/product/domain/model"
	proto "github.com/machinism1011/microservice/product/proto"
)

func SwapTo(request, category interface{}) (err error) {
	data, err := json.Marshal(request)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, category)
}

func ProductToResponseSlice(productSlice []model.Product, response *proto.AllProduct) error {
	for _, cg := range productSlice {
		pi := &proto.ProductInfo{}
		err := SwapTo(cg, pi)
		if err != nil {
			return err
		}
		response.ProductInfo = append(response.ProductInfo, pi)
	}
	return nil
}
