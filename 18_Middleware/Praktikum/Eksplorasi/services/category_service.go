package services

import (
	"rest/database"
	"rest/models"
)

type CategoryService struct{}

func InitCategoryService() CategoryService {
	return CategoryService{}
}

func (service *CategoryService) GetCategories() ([]models.Categories, error) {
	categories := []models.Categories{}

	err := database.DB.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (service *CategoryService) AddCategory(body models.Categories) (models.Categories, error) {
	category := body

	result := database.DB.Create(&category)
	if err := result.Error; err != nil {
		return models.Categories{}, err
	}

	return category, nil
}

func (service *CategoryService) DeleteCategory(id string) error {
	category := models.Categories{}

	result := database.DB.Delete(&category, &id)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}
