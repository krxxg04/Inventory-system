package handlers

import (
	"strconv"

	"inventory-management-api-go/internal/application"
	"inventory-management-api-go/internal/interfaces/http/dto"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *application.CategoryService
}

func NewCategoryHandler(service *application.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(application.NewValidationError("invalid request body"))
		return
	}

	entity, err := h.service.Create(c.Request.Context(), req.Name, req.Description)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, gin.H{"data": dto.ToCategoryResponse(entity)})
}

func (h *CategoryHandler) List(c *gin.Context) {
	entities, err := h.service.List(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"data": dto.ToCategoryResponseList(entities)})
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
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

	c.JSON(200, gin.H{"data": dto.ToCategoryResponse(entity)})
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, parseErr := parseUintParam(c, "id")
	if parseErr != nil {
		c.Error(parseErr)
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(application.NewValidationError("invalid request body"))
		return
	}

	entity, err := h.service.Update(c.Request.Context(), id, req.Name, req.Description)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"data": dto.ToCategoryResponse(entity)})
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, parseErr := parseUintParam(c, "id")
	if parseErr != nil {
		c.Error(parseErr)
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"message": "category deleted successfully"})
}

func parseUintParam(c *gin.Context, key string) (uint, error) {
	raw := c.Param(key)
	id, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return 0, application.NewValidationError("invalid path parameter: " + key)
	}
	return uint(id), nil
}
