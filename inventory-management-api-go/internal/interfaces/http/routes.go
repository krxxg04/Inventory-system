package http

import (
	"net/http"

	"inventory-management-api-go/internal/interfaces/http/handlers"
	"inventory-management-api-go/internal/interfaces/http/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	categoryHandler *handlers.CategoryHandler,
	productHandler *handlers.ProductHandler,
	inventoryHandler *handlers.InventoryHandler,
) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.ErrorMiddleware())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	v1 := router.Group("/api/v1")
	{
		categories := v1.Group("/categories")
		{
			categories.POST("", categoryHandler.Create)
			categories.GET("", categoryHandler.List)
			categories.GET("/:id", categoryHandler.GetByID)
			categories.PUT("/:id", categoryHandler.Update)
			categories.DELETE("/:id", categoryHandler.Delete)
		}

		products := v1.Group("/products")
		{
			products.POST("", productHandler.Create)
			products.GET("", productHandler.List)
			products.GET("/:id", productHandler.GetByID)
			products.PUT("/:id", productHandler.Update)
			products.DELETE("/:id", productHandler.Delete)
			products.GET("/:id/stock", productHandler.GetStock)
			products.GET("/:id/movements", inventoryHandler.ListByProductID)
		}

		inventory := v1.Group("/inventory")
		{
			inventory.POST("/inbound", inventoryHandler.RegisterInbound)
			inventory.POST("/outbound", inventoryHandler.RegisterOutbound)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": gin.H{
				"code":    "NOT_FOUND",
				"message": "route not found",
			},
		})
	})

	return router
}
