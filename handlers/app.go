package handlers

import (
	"final-project-4/database"
	_ "final-project-4/docs"
	"final-project-4/handlers/http_handlers"
	"final-project-4/middlewares"
	"final-project-4/repositories/category_repository/category_pg"
	"final-project-4/repositories/product_repository/product_pg"
	"final-project-4/repositories/transaction_history_repository/transaction_history_pg"
	"final-project-4/repositories/user_repository/user_pg"
	"final-project-4/services"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// @title Toko Belanja API
// @version 1.0
// @description This is a server for Toko Belanja.
// @termsOfService http://swagger.io/terms/
// @contact.name Swagger API Team
// @host localhost:8080
// @BasePath /
func StartApp() *gin.Engine {
	database.StartDB()
	db := database.GetPostgresInstance()

	router := gin.Default()

	userRepo := user_pg.NewUserPG(db)
	userService := services.NewUserService(userRepo)
	userHandler := http_handlers.NewUserHandler(userService)

	// seeding admin with email: admin@gmail.com,
	// password: 123456
	userRepo.SeedingAdmin()

	usersRouter := router.Group("/users")
	{
		usersRouter.POST("/register", userHandler.RegisterUser)
		usersRouter.POST("/login", userHandler.LoginUser)
		usersRouter.PATCH("/topup", middlewares.Authentication(), userHandler.TopupBalance)
	}

	categoryRepo := category_pg.NewCategoryPG(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := http_handlers.NewCategoryHandler(categoryService)

	categoryRouter := router.Group("/category")

	categoryRouter.Use(middlewares.Authentication())
	{
		categoryRouter.POST("/", middlewares.CategoryAuthorization(), categoryHandler.CreateCategory)
		categoryRouter.PATCH("/:categoryId", middlewares.CategoryAuthorization(), categoryHandler.UpdateCategory)
		categoryRouter.GET("/", middlewares.CategoryAuthorization(), categoryHandler.GetAllCategory)
		categoryRouter.DELETE("/:categoryId", middlewares.CategoryAuthorization(), categoryHandler.DeleteCategory)
	}

	productRepo := product_pg.NewProductPG(db)
	productService := services.NewProductService(productRepo)
	productHandler := http_handlers.NewProductHandler(productService)

	productsRouter := router.Group("/products")
	productsRouter.Use(middlewares.Authentication())
	{
		productsRouter.POST("/", middlewares.ProductAuthorization(), productHandler.CreateProduct)
		productsRouter.GET("/", productHandler.GetAllProducts) // fitur ini emg ga dicek admin atau bukan, jd gapake authorization
		productsRouter.PUT("/:productId", middlewares.ProductAuthorization(), productHandler.UpdateProducts)
		productsRouter.DELETE("/:productId", middlewares.ProductAuthorization(), productHandler.DeleteProduct)
	}

	transactionHistoryRepo := transaction_history_pg.NewTransactionHistoryPG(db)
	transactionHistoryService := services.NewTransactionHistoryService(transactionHistoryRepo, productRepo, userRepo, categoryRepo)
	transactionHistoryHandler := http_handlers.NewTransactionHistoryHandler(transactionHistoryService)

	transactionHistoryRouter := router.Group("/transactions")
	transactionHistoryRouter.Use(middlewares.Authentication())
	{
		transactionHistoryRouter.POST("/", transactionHistoryHandler.CreateTransaction)
		transactionHistoryRouter.GET("/my-transactions", middlewares.Authentication(), transactionHistoryHandler.GetTransactionsByUserID)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router

}
