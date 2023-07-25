package handlers

import (
	"final-project-4/database"
	"final-project-4/docs"
	"final-project-4/handlers/http_handlers"
	"final-project-4/middlewares"
	"final-project-4/repositories/category_repository/category_pg"
	"final-project-4/repositories/product_repository/product_pg"
	"final-project-4/repositories/transaction_history_repository/transaction_history_pg"
	"final-project-4/repositories/user_repository/user_pg"
	"final-project-4/services"

	"os"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// const port = ":8080"

func StartApp() {
	database.StartDB()
	db := database.GetPostgresInstance()

	router := gin.Default()

	router.GET("/health-check-fp4", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"appName": "TokoBelanja",
		})
	})

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
		categoryRouter.POST("/", middlewares.AdminAuthorization(), categoryHandler.CreateCategory)
		categoryRouter.PATCH("/:categoryId", middlewares.AdminAuthorization(), categoryHandler.UpdateCategory)
		categoryRouter.GET("/", middlewares.AdminAuthorization(), categoryHandler.GetAllCategory)
		categoryRouter.DELETE("/:categoryId", middlewares.AdminAuthorization(), categoryHandler.DeleteCategory)
	}

	productRepo := product_pg.NewProductPG(db)
	productService := services.NewProductService(productRepo)
	productHandler := http_handlers.NewProductHandler(productService)

	productsRouter := router.Group("/products")
	productsRouter.Use(middlewares.Authentication())
	{
		productsRouter.POST("/", middlewares.AdminAuthorization(), productHandler.CreateProduct)
		productsRouter.GET("/", productHandler.GetAllProducts) // fitur ini emg ga dicek admin atau bukan, jd gapake authorization
		productsRouter.PUT("/:productId", middlewares.AdminAuthorization(), productHandler.UpdateProducts)
		productsRouter.DELETE("/:productId", middlewares.AdminAuthorization(), productHandler.DeleteProduct)
	}

	transactionHistoryRepo := transaction_history_pg.NewTransactionHistoryPG(db, productRepo, userRepo, categoryRepo)
	transactionHistoryService := services.NewTransactionHistoryService(transactionHistoryRepo, productRepo, userRepo, categoryRepo)
	transactionHistoryHandler := http_handlers.NewTransactionHistoryHandler(transactionHistoryService)

	transactionHistoryRouter := router.Group("/transactions")
	transactionHistoryRouter.Use(middlewares.Authentication())
	{
		transactionHistoryRouter.POST("/", transactionHistoryHandler.CreateTransaction)
		transactionHistoryRouter.GET("/my-transactions", transactionHistoryHandler.GetTransactionsByUserID)
		transactionHistoryRouter.GET("/user-transactions", middlewares.AdminAuthorization(), middlewares.Authentication(), transactionHistoryHandler.GetUserTransactions)

	}

	docs.SwaggerInfo.Title = "API Toko Belanja"
	docs.SwaggerInfo.Description = "Ini adalah server API Toko Belanja."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "final-project-4-production.up.railway.app"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}
	// docs.SwaggerInfo.Host = "localhost:8080"
	// docs.SwaggerInfo.Schemes = []string{"http"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":" + os.Getenv("PORT"))

	// router.Run(port)

}
