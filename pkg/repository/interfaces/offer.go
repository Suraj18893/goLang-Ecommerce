package interfaces

import (
	"github.com/Suraj18893/go-Ecommerce/pkg/domain"
	"github.com/Suraj18893/go-Ecommerce/pkg/utils/models"
)

type OfferRepository interface {
	AddNewOffer(models.CreateOffer) error
	MakeOfferExpired(categorytId int) error
	FindDiscountPercentage(categorytId int) (int, error)
	GetOffers(page, limit int) ([]domain.Offer, error)
}
