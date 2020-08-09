package controllers

import (
	"log"
	"smartwallet-api/application/models"
	"smartwallet-api/infrastructure/repositories"

	mapper "github.com/PeteProgrammer/go-automapper"
	"github.com/gin-gonic/gin"
)

type MarketDataController struct {
	MarketDataRepository repositories.MarketDataRepository
}

func NewMarketDataController(m repositories.MarketDataRepository) MarketDataController {
	return MarketDataController{m}
}

func (controller *MarketDataController) GetAll(c *gin.Context) {
	var model []models.MarketData
	dtos := controller.MarketDataRepository.GetAll()
	mapper.Map(dtos, &model)

	c.JSON(200, model)
}

type MarketData struct {
	ID string `uri:"id" binding:"required"`
}

func (controller *MarketDataController) GetById(c *gin.Context) {
	var marketData MarketData
	if err := c.ShouldBindUri(&marketData); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	log.Println("ID: " + marketData.ID)

	dto := controller.MarketDataRepository.GetById(marketData.ID)
	c.JSON(200, dto)
}
