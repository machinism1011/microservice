package service

import (
	"github.com/machinism1011/microservice/category/domain/model"
	"github.com/machinism1011/microservice/category/domain/repository"
)

type ICategoryDataService interface {
	AddCategory(*model.Category) (int64, error)
	DeleteCategory(int64) error
	UpdateCategory(*model.Category) error
	FindCategoryByID(int64) (*model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByLevel(uint322 uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
	FindAllCategory() ([]model.Category, error)
}

type CategoryDataService struct {
	CategoryRepository repository.ICategoryRepository
}

func NewCategoryDataService(categoryRepository repository.ICategoryRepository) ICategoryDataService{
	return &CategoryDataService{ categoryRepository }
}

func (c *CategoryDataService) AddCategory(category *model.Category) (int64, error) {
	return c.CategoryRepository.CreateCategory(category)
}

func (c *CategoryDataService) DeleteCategory(categoryID int64) error {
	return c.CategoryRepository.DeleteCategoryByID(categoryID)
}

func (c *CategoryDataService) UpdateCategory(category *model.Category) error {
	return c.CategoryRepository.UpdateCategory(category)
}

func (c *CategoryDataService) FindCategoryByID(categoryID int64) (*model.Category, error) {
	return c.CategoryRepository.FindCategoryByID(categoryID)
}

func (c *CategoryDataService) FindCategoryByName(categoryName string) (*model.Category, error) {
	return c.CategoryRepository.FindCategoryByName(categoryName)
}

func (c *CategoryDataService) FindCategoryByLevel(level uint32) ([]model.Category, error) {
	return c.CategoryRepository.FindCategoryByLevel(level)
}

func (c *CategoryDataService) FindCategoryByParent(parent int64) ([]model.Category, error) {
	return c.CategoryRepository.FindCategoryByParent(parent)
}

func (c *CategoryDataService) FindAllCategory() ([]model.Category, error) {
	return c.CategoryRepository.FindAllCategory()
}