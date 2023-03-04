package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"ProjectBuahIn/buah"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type buahHandler struct {
	buahService buah.Service
}

func NewBuahHandler(buahService buah.Service) *buahHandler {
	return &buahHandler{buahService}
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


func (h *buahHandler) QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

func (h *buahHandler) GetBuah(c *gin.Context) {
	idString := c.Param("id")

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

func convertToBuahResponse(b buah.Buah) buah.BuahResponse {
	return buah.BuahResponse{
		ID:          b.ID,
		Jenis:       b.Jenis,
		Price:       b.Price,
		Description: b.Description,
		Matang:      b.Matang,
		Discount:    b.Discount,
	}
}
