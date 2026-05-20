package handlers

import (
	"inventory-management-api-go/internal/application"
	"inventory-management-api-go/internal/interfaces/http/dto"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *application.ProductService
}

func NewProductHandler(service *application.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(application.NewValidationError("invalid request body"))
		return
	}

	entity, err := h.service.Create(
		c.Request.Context(),
		req.Name,
		req.Description,
		req.SKU,
		req.Price,
		req.Stock,
		req.CategoryID,
	)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, gin.H{"data": dto.ToProductResponse(entity)})
}

func (h *ProductHandler) List(c *gin.Context) {
	entities, err := h.service.List(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"data": dto.ToProductResponseList(entities)})
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	id, parseErr := parseUintParam(c, "id")
	if parseErr != nil {
		c.Error(parseErr)
		return
	}

	entity, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"data": dto.ToProductResponse(entity)})
}

func (h *ProductHandler) Update(c *gin.Context) {
	id, parseErr := parseUintParam(c, "id")
	if parseErr != nil {
		c.Error(parseErr)
		return
	}

	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(application.NewValidationError("invalid request body"))
		return
	}

	entity, err := h.service.Update(
		c.Request.Context(),
		id,
		req.Name,
		req.Description,
		req.SKU,
		req.Price,
		req.CategoryID,
	)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"data": dto.ToProductResponse(entity)})
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id, parseErr := parseUintParam(c, "id")
	if parseErr != nil {
		c.Error(parseErr)
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"message": "product deleted successfully"})
}

func (h *ProductHandler) GetStock(c *gin.Context) {
	id, parseErr := parseUintParam(c, "id")
	if parseErr != nil {
		c.Error(parseErr)
		return
	}

	stock, err := h.service.GetStock(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{
		"data": dto.ProductStockResponse{
			ProductID: id,
			Stock:     stock,
		},
	})
}
