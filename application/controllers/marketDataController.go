package controllers

import (
	mapper "github.com/PeteProgrammer/go-automapper"
	"github.com/gin-gonic/gin"
	"smartwallet-api/application/models"
	"smartwallet-api/infrastructure/repositories"
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
