package web

import (
	"fmt"

	"github.com/AasSuhendar/go-restful-api/models/domain"
)

type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ToCategoryResponse(category domain.Category) CategoryResponse {
	return CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []CategoryResponse {
	var categoryResponses []CategoryResponse
	fmt.Println("helper", categories)
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
