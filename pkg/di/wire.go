//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Suraj18893/go-Ecommerce/pkg/api"
	"github.com/Suraj18893/go-Ecommerce/pkg/api/handlers"
	"github.com/Suraj18893/go-Ecommerce/pkg/config"
	"github.com/Suraj18893/go-Ecommerce/pkg/db"
	"github.com/Suraj18893/go-Ecommerce/pkg/repository"
	"github.com/Suraj18893/go-Ecommerce/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
	wire.Build(db.ConnectDB,
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		handlers.NewUserHandler,

		repository.NewAdminRepository,
		usecase.NewAdminUsecase,
		handlers.NewAdminHandler,

		repository.NewCartRepository,
		usecase.NewCartUsecase,
		handlers.NewCartHandler,

		repository.NewCategoryRepository,
		usecase.NewCategoryUsecase,
		handlers.NewCategoryHandler,

		repository.NewInventoryRepository,
		usecase.NewInventoryUsecase,
		handlers.NewInventoryHandler,

		repository.NewOfferRepository,
		usecase.NewOfferUsecase,
		handlers.NewOfferHandler,

		repository.NewOtpRepository,
		usecase.NewOtpUsecase,
		handlers.NewOtpHandler,

		repository.NewCouponRepository,
		usecase.NewCouponUsecase,
		handlers.NewCouponHandler,

		repository.NewPaymentRepository,
		usecase.NewPaymentUsecase,
		handlers.NewPaymentHandler,

		repository.NewWishlistRepository,
		usecase.NewWishlistUsecase,
		handlers.NewWishlistHandler,

		repository.NewOrderRepository,
		usecase.NewOrderUsecase,
		handlers.NewOrderHandler,

		api.NewServerHttp)
	return &api.ServerHTTP{}, nil
}
