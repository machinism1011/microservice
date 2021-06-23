package handler

import (
	"context"
	"github.com/machinism1011/microservice/category/common"
	"github.com/machinism1011/microservice/category/domain/model"
	"github.com/machinism1011/microservice/category/domain/service"
	proto "github.com/machinism1011/microservice/category/proto/category"
)

type Category struct{
	CategoryDataService service.ICategoryDataService
}

func (c *Category) CreateCategory(_ context.Context, request *proto.CategoryRequest, response *proto.CreateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(request, category)
	if err != nil {
		return err
	}

	categoryID, err := c.CategoryDataService.AddCategory(category)
	if err != nil {
		return err
	}
	response.Message = "分类添加成功"
	response.CategoryId = categoryID
	return nil
}

func (c *Category) UpdateCategory(_ context.Context, request *proto.CategoryRequest, response *proto.UpdateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(request, category)
	if err != nil {
		return err
	}

	err = c.CategoryDataService.UpdateCategory(category)
	if err != nil {
		return err
	}
	response.Message = "分类更新成功"
	return nil
}

func (c *Category) DeleteCategory(_ context.Context, request *proto.DeleteCategoryRequest, response *proto.DeleteCategoryResponse) error {
	err := c.CategoryDataService.DeleteCategory(request.CategoryId)
	if err != nil {
		return err
	}
	response.Message = "分类删除成功"
	return nil
}

func (c *Category) FindCategoryByName(_ context.Context, request *proto.FindByNameRequest, response *proto.CategoryResponse) error {
	category, err := c.CategoryDataService.FindCategoryByName(request.CategoryName)
	if err != nil {
		return err
	}
	return common.SwapTo(category, response)
}

func (c *Category) FindCategoryByID(_ context.Context, request *proto.FindByIDRequest, response *proto.CategoryResponse) error{
	category, err := c.CategoryDataService.FindCategoryByID(request.CategoryId)
	if err != nil {
		return err
	}
	return common.SwapTo(category, response)
}

func (c *Category) FindCategoryByLevel(_ context.Context, request *proto.FindByLevelRequest, response *proto.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByLevel(request.CategoryLevel)
	if err != nil {
		return err
	}
	return common.CategoryToResponseSlice(categorySlice, response)
}
func (c *Category) FindCategoryByParent(_ context.Context, request *proto.FindByParentRequest, response *proto.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByParent(request.CategoryParent)
	if err != nil {
		return err
	}
	return common.CategoryToResponseSlice(categorySlice, response)
}

func (c *Category) FindAllCategory(_ context.Context, _ *proto.FindAllRequest, response *proto.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return err
	}
	return common.CategoryToResponseSlice(categorySlice, response)
}