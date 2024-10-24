package handlers

import (
	"net/http"
	"strconv"

	services "github.com/Suraj18893/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Suraj18893/go-Ecommerce/pkg/utils/models"
	"github.com/Suraj18893/go-Ecommerce/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type OfferHandler struct {
	offerUsecase services.OfferUsecase
}

// Constructor function

func NewOfferHandler(offerUsecase services.OfferUsecase) *OfferHandler {
	return &OfferHandler{
		offerUsecase: offerUsecase,
	}
}

// @Summary		Add Offer
// @Description	Admin can add new  offers
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param           offer      body     models.CreateOffer   true   "Offer"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/offers/create [post]
func (offH *OfferHandler) AddOffer(c *gin.Context) {
	var offer models.CreateOffer

	if err := c.BindJSON(&offer); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := offH.offerUsecase.AddNewOffer(offer); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "offer adding failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "offer added successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Expire Offer
// @Description	Admin can add Expire  offers
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param           catID      query     string   true   "Category ID"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/offers/expire [post]
func (offH *OfferHandler) ExpireValidity(c *gin.Context) {
	catIdStr := c.Query("catID")
	catId, err := strconv.Atoi(catIdStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := offH.offerUsecase.MakeOfferExpire(catId); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "make offer expired failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully turned the offer invalid", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		List Offers
// @Description	Admin can view the list of  offers
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/offers [get]
func (offH *OfferHandler) Offers(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	offers, err := offH.offerUsecase.GetOffers(page, limit)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "couldn't get offers", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully retrieved the offers", offers, nil)
	c.JSON(http.StatusOK, successRes)

}
