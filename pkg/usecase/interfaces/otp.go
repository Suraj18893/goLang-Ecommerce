package interfaces

import "github.com/Suraj18893/go-Ecommerce/pkg/utils/models"

type OtpUsecase interface {
	VerifyOTP(code models.VerifyData) (models.UserToken, error)
	SendOTP(phone string) error
}
