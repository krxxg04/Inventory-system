package main

import (
	"log"
	"os"

	"inventory-management-api-go/internal/application"
	"inventory-management-api-go/internal/infrastructure/database"
	"inventory-management-api-go/internal/infrastructure/persistence"
	httpapi "inventory-management-api-go/internal/interfaces/http"
	"inventory-management-api-go/internal/interfaces/http/handlers"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	dbConfig := database.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		Name:     getEnv("DB_NAME", "inventory_db"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	db, err := database.NewPostgresConnection(dbConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("failed to run database migrations: %v", err)
	}

	categoryRepo := persistence.NewCategoryRepositoryGorm(db)
	productRepo := persistence.NewProductRepositoryGorm(db)
	inventoryRepo := persistence.NewInventoryRepositoryGorm(db)
	txManager := database.NewGormTxManager(db)

	categoryService := application.NewCategoryService(categoryRepo)
	productService := application.NewProductService(productRepo, categoryRepo, inventoryRepo)
	inventoryService := application.NewInventoryService(productRepo, inventoryRepo, txManager)

	categoryHandler := handlers.NewCategoryHandler(categoryService)
	productHandler := handlers.NewProductHandler(productService)
	inventoryHandler := handlers.NewInventoryHandler(inventoryService)

	router := httpapi.NewRouter(categoryHandler, productHandler, inventoryHandler)
	port := getEnv("APP_PORT", "8080")

	log.Printf("server started on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
