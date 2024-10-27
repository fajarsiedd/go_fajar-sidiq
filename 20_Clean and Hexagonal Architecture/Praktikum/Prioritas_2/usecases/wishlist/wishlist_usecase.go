package wishlist

import (
	"go-wishlist-api/handlers/wishlist/request"
	"go-wishlist-api/models"
)

type WishlistUsecase interface {
	GetAll() ([]models.Wishlist, error)
	Create(input request.WishlistInput) (models.Wishlist, error)
}
