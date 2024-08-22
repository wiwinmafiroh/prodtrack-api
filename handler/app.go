package handler

import (
	"fmt"
	"os"
	_ "prodtrack-api/docs"

	"prodtrack-api/database"
	"prodtrack-api/repository/product_repository/product_postgres"
	"prodtrack-api/repository/user_repository/user_postgres"
	"prodtrack-api/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title The ProdTrack API Documentation
// @version 1.0
// @description ProdTrack API is a comprehensive solution for managing product data. It includes full CRUD operations, along with robust authentication and authorization mechanisms. Admin have the ability to manage all products, while regular users are restricted to managing only their own products. The API is built using Golang and the Gin framework.
// @contact.name Wiwin Mafiroh
// @contact.email wiwinmafiroh@gmail.com
// @host localhost:8080
// @BasePath /
// @schemes http
func StartApp() {
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))

	db := database.GetDatabaseInstance()
	defer db.Close()

	userRepository := user_postgres.NewUserPostgres(db)
	userService := service.NewUserService(userRepository)
	userHandler := NewUserHandler(userService)

	productRepository := product_postgres.NewProductPostgres(db)
	productService := service.NewProductService(productRepository)
	productHandler := NewProductHandler(productService)

	authService := service.NewAuthService(userRepository, productRepository)

	route := gin.Default()

	route.GET("/", Welcome)
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.UserRegister)
		userRoute.POST("/login", userHandler.UserLogin)
	}

	productRoute := route.Group("/products")
	{
		productRoute.Use(authService.Authentication())
		productRoute.Use(authService.AuthorizationRole())
		productRoute.POST("/", productHandler.CreateProduct)
		productRoute.GET("/", productHandler.GetProducts)
		productRoute.Use(authService.AuthorizationProduct())
		productRoute.GET("/:productId", productHandler.GetProductById)
		productRoute.PUT("/:productId", productHandler.UpdateProductById)
		productRoute.DELETE("/:productId", productHandler.DeleteProductById)
	}

	route.Run(PORT)
}
