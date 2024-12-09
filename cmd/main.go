package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "leadgen-test-task/docs"
	"leadgen-test-task/internal/handlers"
	"leadgen-test-task/internal/repository"
	"leadgen-test-task/pkg/database"
	"leadgen-test-task/pkg/logs"
)

func main() {
	//logger init
	logger := logs.NewLogger()

	// Setup config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf("Error reading config file: %v", err)
	}

	// Database config
	dbConfig := database.Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.name"),
	}

	// Establish DB connection
	db, err := database.NewPostgresConnection(dbConfig)
	if err != nil {
		logger.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	if err := database.InitDatabase(db); err != nil {
		logger.Fatalf("Could not initialize database: %v", err)
	}

	// Setup repositories and handlers
	buildingRepo := repository.NewBuildingRepository(db)
	buildingHandler := handlers.NewBuildingHandler(buildingRepo, logger)

	// Setup Gin router
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/buildings", buildingHandler.CreateBuilding)
		v1.GET("/buildings", buildingHandler.ListBuildings)
	}

	// Swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}

	logger.Infof("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
