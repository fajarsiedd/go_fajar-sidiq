package wishlist

import (
	"errors"
	"go-wishlist-api/handlers/wishlist/request"
	"go-wishlist-api/models"
	"go-wishlist-api/repositories/wishlist"
)

type wishlistUsecase struct {
	repository wishlist.WishlistRepository
}

func NewWishlistUsecase(r wishlist.WishlistRepository) *wishlistUsecase {
	return &wishlistUsecase{
		repository: r,
	}
}

func (usecase wishlistUsecase) GetAll() ([]models.Wishlist, error) {
	return usecase.repository.GetAll()
}

func (usecase wishlistUsecase) Create(input request.WishlistInput) (models.Wishlist, error) {
	if input.Title == "" {
		return models.Wishlist{}, errors.New("title cannot be empty")
	}

	return usecase.repository.Create(input)
}
