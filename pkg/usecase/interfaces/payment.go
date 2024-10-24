package interfaces

import (
	// "github.com/Suraj18893/go-Ecommerce/pkg/domain"
	"github.com/Suraj18893/go-Ecommerce/pkg/utils/models"
)

type PaymentUsecase interface {
	AddNewPaymentMethod(paymentMethod string) error
	RemovePaymentMethod(paymentMethodID int) error
	GetPaymentMethods() ([]models.PaymentMethod, error)
	MakePaymentRazorPay(orderID string, userID int) (models.OrderPaymentDetails, error)
	VerifyPayment(paymentID string, razorID string, orderID string) error
}
