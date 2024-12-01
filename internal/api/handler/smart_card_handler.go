package handler

import (
	"strconv"

	"github.com/adiubaidah/rfid-syafiiyah/internal/constant/model"
	db "github.com/adiubaidah/rfid-syafiiyah/internal/storage/persistence"
	"github.com/adiubaidah/rfid-syafiiyah/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SmartCardHandler interface {
	ListSmartCardsHandler(c *gin.Context)
	UpdateSmartCardHandler(c *gin.Context)
	DeleteSmartCardHandler(c *gin.Context)
}

type smartCardHandler struct {
	logger  *logrus.Logger
	usecase usecase.SmartCardUseCase
}

func NewSmartCardHandler(logger *logrus.Logger, usecase usecase.SmartCardUseCase) SmartCardHandler {
	return &smartCardHandler{logger: logger, usecase: usecase}
}

func (h *smartCardHandler) ListSmartCardsHandler(c *gin.Context) {
	var request model.ListSmartCardRequest
	if err := c.ShouldBind(&request); err != nil {
		h.logger.Error(err)
		c.JSON(400, model.ResponseMessage{Code: 400, Status: "error", Message: err.Error()})
		return
	}

	if request.Limit == 0 {
		request.Limit = 10
	}

	if request.Page == 0 {
		request.Page = 1
	}

	if request.CardOwner == "" {
		request.CardOwner = db.CardOwnerAll
	}

	result, err := h.usecase.ListSmartCards(c, &request)
	if err != nil {
		h.logger.Error("error lur", err)
		c.JSON(500, model.ResponseMessage{Code: 500, Status: "error", Message: err.Error()})
		return
	}

	count, err := h.usecase.CountSmartCards(c, &request)
	if err != nil {
		h.logger.Error(err)
		c.JSON(500, model.ResponseMessage{Code: 500, Status: "error", Message: err.Error()})
		return
	}

	pagination := model.Pagination{
		CurrentPage:  request.Page,
		TotalPages:   int32((count + int64(request.Limit) - 1) / int64(request.Limit)),
		TotalItems:   count,
		ItemsPerPage: request.Limit,
	}

	c.JSON(200, model.ResponseData[model.ListSmartCardResponse]{
		Code:   200,
		Status: "success",
		Data: model.ListSmartCardResponse{
			Items:      *result,
			Pagination: pagination,
		},
	})
}

func (h *smartCardHandler) UpdateSmartCardHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		c.JSON(400, model.ResponseMessage{Code: 400, Status: "error", Message: "Invalid ID"})
		return
	}
	smartCardId := int32(id)

	var smartCardRequest model.UpdateSmartCardRequest
	if err := c.ShouldBind(&smartCardRequest); err != nil {
		h.logger.Error(err)
		c.JSON(400, model.ResponseMessage{Code: 400, Status: "error", Message: err.Error()})
		return
	}

	result, err := h.usecase.UpdateSmartCard(c, &smartCardRequest, smartCardId)
	if err != nil {
		h.logger.Error(err)
		c.JSON(500, model.ResponseMessage{Code: 500, Status: "error", Message: err.Error()})
		return
	}

	c.JSON(200, model.ResponseData[model.SmartCardComplete]{Code: 200, Status: "success", Data: *result})
}

func (h *smartCardHandler) DeleteSmartCardHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		c.JSON(400, model.ResponseMessage{Code: 400, Status: "error", Message: "Invalid ID"})
		return
	}
	smartCardId := int32(id)

	deletedSmartCard, err := h.usecase.DeleteSmartCard(c, smartCardId)
	if err != nil {
		h.logger.Error(err)
		c.JSON(500, model.ResponseMessage{Code: 500, Status: "error", Message: err.Error()})
		return
	}

	c.JSON(200, model.ResponseData[model.SmartCard]{Code: 200, Status: "success", Data: *deletedSmartCard})
}