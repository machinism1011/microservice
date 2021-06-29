package handler

import (
	"context"
	"github.com/machinism1011/microservice/product/common"
	"github.com/machinism1011/microservice/product/domain/model"
	"github.com/machinism1011/microservice/product/domain/service"
	protoProduct "github.com/machinism1011/microservice/product/proto"
)

type Product struct{
	ProductDataService service.IProductDataService
}


func (p *Product) AddProduct(_ context.Context, request *protoProduct.ProductInfo, response *protoProduct.ResponseProduct) error {
	product := &model.Product{}
	if err := common.SwapTo(request, product); err != nil {
		return err
	}
	productID, err := p.ProductDataService.AddProduct(product)
	if err != nil {
		return err
	}
	response.ProductId = productID
	return nil
}

func (p *Product) FindProductByID(_ context.Context, request *protoProduct.RequestID,  response *protoProduct.ProductInfo) error {
	productData, err := p.ProductDataService.FindProductByID(request.ProductId)
	if err != nil {
		return err
	}
	if err := common.SwapTo(productData, response); err != nil {
		return err
	}
	return nil
}
func (p *Product) UpdateProduct(_ context.Context, request *protoProduct.ProductInfo, response *protoProduct.ResponseMessage) error {
	product := &model.Product{}
	if err := common.SwapTo(request, product); err != nil {
		return err
	}
	err := p.ProductDataService.UpdateProduct(product)
	if err != nil {
		return err
	}
	response.Message = "更新成功"
	return nil
}

func (p *Product) DeleteProductByID(_ context.Context, request *protoProduct.RequestID, response *protoProduct.ResponseMessage) error {
	if err := p.ProductDataService.DeleteProduct(request.ProductId); err != nil {
		return err
	}
	response.Message = "删除成功"
	return nil
}
func (p *Product) FindAllProduct(_ context.Context, _ *protoProduct.RequestAll, response *protoProduct.AllProduct) error {
	productSlice, err := p.ProductDataService.FindAllProduct()
	if err != nil {
		return err
	}

	if err := common.ProductToResponseSlice(productSlice, response); err != nil {
		return err
	}
	return nil
}

