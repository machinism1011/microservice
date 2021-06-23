package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/machinism1011/microservice/category/domain/model"
)

type ICategoryRepository interface {
	InitTable() error
	FindCategoryByID(int64) (*model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
	FindAllCategory() ([]model.Category, error)
	CreateCategory(category *model.Category) (int64, error)
	DeleteCategoryByID(int64) error
	UpdateCategory(*model.Category) error
}

type CategoryRepository struct {
	mysqlDb *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{mysqlDb:db}
}

func (c *CategoryRepository) InitTable() error {
	return c.mysqlDb.CreateTable(&model.Category{}).Error
}

func (c *CategoryRepository) FindCategoryByID(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	return category, c.mysqlDb.First(category, id).Error
}

func (c *CategoryRepository) FindCategoryByName(categoryName string) (category *model.Category, err error) {
	category = &model.Category{}
	return category, c.mysqlDb.Where("category_name = ?", categoryName).Find(category).Error
}

func (c *CategoryRepository) FindCategoryByLevel(level uint32) (categorySlice []model.Category, err error) {
	categorySlice = make([]model.Category, 0)
	return categorySlice, c.mysqlDb.Where("category_level = ?", level).Find(categorySlice).Error
}

func (c *CategoryRepository) FindCategoryByParent(parentID int64) (categorySlice []model.Category, err error) {
	categorySlice = make([]model.Category, 0)
	return categorySlice, c.mysqlDb.Where("category_parent = ?", parentID).Find(categorySlice).Error
}

func (c *CategoryRepository) FindAllCategory() (categorySlice []model.Category, err error) {
	categorySlice = make([]model.Category, 0)
	return categorySlice, c.mysqlDb.Find(&categorySlice).Error
}

func (c *CategoryRepository) CreateCategory(category *model.Category) (categoryID int64, err error) {
	return category.ID, c.mysqlDb.Create(category).Error
}

func (c *CategoryRepository) DeleteCategoryByID(id int64) error {
	return c.mysqlDb.Where("id = ?", id).Delete(&model.Category{}).Error
}

func (c *CategoryRepository) UpdateCategory(category *model.Category) error {
	return c.mysqlDb.Model(category).Update(category).Error
}

