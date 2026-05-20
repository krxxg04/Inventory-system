package handlers

import (
	"inventory-management-api-go/internal/application"
	"inventory-management-api-go/internal/interfaces/http/dto"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	service *application.InventoryService
}

func NewInventoryHandler(service *application.InventoryService) *InventoryHandler {
	return &InventoryHandler{service: service}
}

func (h *InventoryHandler) RegisterInbound(c *gin.Context) {
	var req dto.InventoryMovementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(application.NewValidationError("invalid request body"))
		return
	}

	entity, err := h.service.RegisterInbound(
		c.Request.Context(),
		req.ProductID,
		req.Quantity,
		req.Description,
	)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, gin.H{"data": dto.ToInventoryMovementResponse(entity)})
}

func (h *InventoryHandler) RegisterOutbound(c *gin.Context) {
	var req dto.InventoryMovementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(application.NewValidationError("invalid request body"))
		return
	}

	entity, err := h.service.RegisterOutbound(
		c.Request.Context(),
		req.ProductID,
		req.Quantity,
		req.Description,
	)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, gin.H{"data": dto.ToInventoryMovementResponse(entity)})
}

func (h *InventoryHandler) ListByProductID(c *gin.Context) {
	id, parseErr := parseUintParam(c, "id")
	if parseErr != nil {
		c.Error(parseErr)
		return
	}

	entities, err := h.service.ListMovementsByProductID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"data": dto.ToInventoryMovementResponseList(entities)})
}
