package common

import (
	"encoding/json"
	"github.com/machinism1011/microservice/category/domain/model"
	proto "github.com/machinism1011/microservice/category/proto/category"
)

// 通过json tag进行结构体赋值
func SwapTo(request, category interface{}) (err error) {
	data, err := json.Marshal(request)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, category)
}

func CategoryToResponseSlice(categorySlice []model.Category, response *proto.FindAllResponse) error {
	for _, cg := range categorySlice {
		cr := &proto.CategoryResponse{}
		err := SwapTo(cg, cr)
		if err != nil {
			return err
		}
		response.Category = append(response.Category, cr)
	}
	return nil
}