package service

import (
	"context"

	"github.com/jonathanvinicius/grpc_full_cycle/internal/database"
	"github.com/jonathanvinicius/grpc_full_cycle/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:						  category.ID,
		Name: 					category.Name,
		Description:    category.Description,
	}

	return categoryResponse, nil
}