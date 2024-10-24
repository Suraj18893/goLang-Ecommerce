package interfaces

import "github.com/Suraj18893/go-Ecommerce/pkg/utils/models"

type OtpRepository interface {
	FindUserByMobileNumber(phone string)bool
	UserDetailsUsingPhone(phone string)(models.UserDetailsResponse,error)
}