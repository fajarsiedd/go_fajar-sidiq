package services

import (
	"rest/database"
	"rest/models"
)

type PostService struct{}

func InitPostService() PostService {
	return PostService{}
}

func (service *PostService) GetPosts() ([]models.Posts, error) {
	posts := []models.Posts{}

	err := database.DB.Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (service *PostService) GetPostByID(id string) (models.Posts, error) {
	post := models.Posts{}

	err := database.DB.First(&post, &id).Error
	if err != nil {
		return models.Posts{}, err
	}

	return post, nil
}

func (service *PostService) AddPost(body models.Posts) (models.Posts, error) {
	post := body

	result := database.DB.Create(&post)
	if err := result.Error; err != nil {
		return models.Posts{}, err
	}

	return post, nil
}

func (service *PostService) UpdatePost(body models.Posts) (models.Posts, error) {
	post := body

	result := database.DB.Updates(&post)
	if err := result.Error; err != nil {
		return models.Posts{}, err
	}

	return post, nil
}

func (service *PostService) DeletePost(id string) error {
	post := models.Posts{}

	result := database.DB.Delete(&post, &id)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}
