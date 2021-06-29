package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/machinism1011/microservice/product/domain/model"
)

type IProductRepository interface {
	InitTable() error
	FindProductByID(int64) (*model.Product, error)
	FindAllProduct() ([]model.Product, error)
	CreateProduct(*model.Product) (int64, error)
	DeleteProductByID(int64) error
	UpdateProduct(*model.Product) error
}

type ProductRepository struct {
	mysqlDb *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{mysqlDb: db}
}

func (p *ProductRepository) InitTable() error {
	// create multi tables
	return p.mysqlDb.CreateTable(&model.Product{}, &model.ProductImage{}, &model.ProductSeo{}, &model.ProductSize{}).Error
}

func (p *ProductRepository) FindProductByID(productID int64) (product *model.Product, err error) {
	product = &model.Product{}
	// 加载关联信息，使用preload
	return product, p.mysqlDb.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").First(product, productID).Error
}

func (p *ProductRepository) FindAllProduct() (productSlice []model.Product, err error) {
	productSlice = make([]model.Product, 0)
	return productSlice, p.mysqlDb.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").Find(&productSlice).Error
}

func (p *ProductRepository) CreateProduct(product *model.Product) (productID int64, err error) {
	return product.ID, p.mysqlDb.Create(product).Error
}

func (p *ProductRepository) DeleteProductByID(productID int64) error {
	// 关联表，需要开启事务
	tx := p.mysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Unscoped().Where("id = ?", productID).Delete(&model.Product{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Where("images_product_id = ?", productID).Delete(&model.ProductImage{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Where("size_product_id = ?", productID).Delete(&model.ProductSize{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Where("seo_product_id = ?", productID).Delete(&model.ProductSeo{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (p *ProductRepository) UpdateProduct(product *model.Product) error {
	return p.mysqlDb.Model(product).Update(product).Error
}
