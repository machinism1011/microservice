package service

import (
	"github.com/machinism1011/microservice/cart/domain/model"
	"github.com/machinism1011/microservice/cart/domain/repository"
)

type IProductDataService interface {
	AddProduct(*model.Product) (int64, error)
	DeleteProduct(int64) error
	UpdateProduct(*model.Product) error
	FindProductByID(int64) (*model.Product, error)
	FindAllProduct() ([]model.Product, error)
}

type ProductDataService struct {
	ProductRepository repository.IProductRepository
}

func NewProductDataService(productRepository repository.IProductRepository) IProductDataService {
	return &ProductDataService{productRepository}
}

func (p *ProductDataService) AddProduct(product *model.Product) (int64, error) {
	return p.ProductRepository.CreateProduct(product)
}

func (p *ProductDataService) DeleteProduct(productID int64) error {
	return p.ProductRepository.DeleteProductByID(productID)
}

func (p *ProductDataService) UpdateProduct(product *model.Product) error {
	return p.ProductRepository.UpdateProduct(product)
}

func (p *ProductDataService) FindProductByID(productID int64) (*model.Product, error) {
	return p.ProductRepository.FindProductByID(productID)
}

func (p *ProductDataService) FindAllProduct() ([]model.Product, error) {
	return p.ProductRepository.FindAllProduct()
}
