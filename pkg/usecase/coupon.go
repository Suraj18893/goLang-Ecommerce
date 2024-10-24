package usecase

import (
	"errors"

	"github.com/Suraj18893/go-Ecommerce/pkg/domain"
	interfaces "github.com/Suraj18893/go-Ecommerce/pkg/repository/interfaces"
	services"github.com/Suraj18893/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Suraj18893/go-Ecommerce/pkg/utils/models"
)

type couponUsecase struct {
	couponRepo interfaces.CouponRepository
}

// constructor function

func NewCouponUsecase(couponRepo interfaces.CouponRepository) services.CouponUsecase {
	return &couponUsecase{
		couponRepo: couponRepo,
	}
}

func (coupU *couponUsecase) Addcoupon(coupon models.Coupon) error {
	if err := coupU.couponRepo.AddCoupon(coupon); err != nil {
		return errors.New("coupon adding failed")
	}
	return nil
}

func (coupU *couponUsecase) MakeCouponInvalid(id int) error {
	if err := coupU.couponRepo.MakeCouponInvalid(id); err != nil {
		return err
	}
	return nil
}

func (coupU *couponUsecase) GetCoupons() ([]domain.Coupon, error) {
	coupons, err := coupU.couponRepo.GetCoupons()
	if err != nil {
		return []domain.Coupon{}, err
	}
	return coupons, nil
}
