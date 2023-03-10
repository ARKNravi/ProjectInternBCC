package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"ProjectBuahIn/buah"
	"ProjectBuahIn/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type buahHandler struct {
	buahService buah.Service
}

func NewBuahHandler(buahService buah.Service) *buahHandler {
	return &buahHandler{buahService}
}

type OrderHandler interface {
	OrderProduct(*gin.Context)
}

type orderHandler struct {
	repo buah.OrderRepository
}

// NewOrderHandler --> return new Order Handler
func NewOrderHandler() OrderHandler {
	return &orderHandler{
		repo: buah.NewOrderRepository(),
	}
}

func (h *buahHandler) GetBuahs(c *gin.Context) {
	buahs, err := h.buahService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var buahsResponse []buah.BuahResponse

	for _, b := range buahs {
		buahResponse := convertToBuahResponse(b)

		buahsResponse = append(buahsResponse, buahResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": buahsResponse,
	})
}

func (h *buahHandler) GetUser(c *gin.Context) {
	idString := c.Param("id")

	id, _ := strconv.Atoi(idString)

	users, err := h.buahService.FindByIDuser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	user := convertToUserResponse(users)
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *buahHandler) GetBuah(c *gin.Context) {
	idString := c.Param("nama")

	id, _ := strconv.Atoi(idString)

	b, err := h.buahService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	buahResponse := convertToBuahResponse(b)
	c.JSON(http.StatusOK, gin.H{
		"data": buahResponse,
	})

}
func (h *buahHandler) GetNamaBuah(c *gin.Context) {
	idString := c.Query("nama")

	b, err := h.buahService.FindByNama(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	buahResponse := convertToBuahResponse(b)
	c.JSON(http.StatusOK, gin.H{
		"data": buahResponse,
	})

}

func (h *buahHandler) CreateBuah(c *gin.Context) {
	//title, price
	var buahRequest buah.BuahRequest

	err := c.ShouldBindJSON(&buahRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	buah, err := h.buahService.Create(buahRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": buah,
	})
}

func (h *buahHandler) UpdateBuah(c *gin.Context) {
	//title, price
	var buahRequest buah.BuahRequest

	err := c.ShouldBindJSON(&buahRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	buah, err := h.buahService.Update(id, buahRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": buah,
	})
}

func (h *buahHandler) DeleteBuah(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	b, err := h.buahService.Delete(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	buahResponse := convertToBuahResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": buahResponse,
	})
}

func (h *buahHandler) OrderBuah(c *gin.Context) {

}

func convertToBuahResponse(b buah.Buah) buah.BuahResponse {
	return buah.BuahResponse{
		Nama:        b.Nama,
		Price:       b.Price,
		Description: b.Description,
		Quantity:    b.Quantity,
		Discount:    b.Discount,
	}
}

func convertToUserResponse(b models.User) models.User {
	return models.User{
		UserID:   b.UserID,
		Username: b.Username,
		Email:    b.Email,
		Password: b.Password,
	}
}

func (h *orderHandler) OrderProduct(ctx *gin.Context) {
	prodIDStr := ctx.Query("product")
	if prodID, err := strconv.Atoi(prodIDStr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		quantityIDStr := ctx.Param("quantity")
		if quantityID, err := strconv.Atoi(quantityIDStr); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			userID := ctx.GetFloat64("userID")
			if err := h.repo.OrderProduct(int(userID), prodID, quantityID); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.String(http.StatusOK, "Product Successfully ordered")
			}
		}
	}

}
