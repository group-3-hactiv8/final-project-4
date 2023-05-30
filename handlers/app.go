package handlers

import (
	"final-project-4/database"
	// _ "final-project-4/docs"
	"final-project-4/handlers/http_handlers"
	"final-project-4/middlewares"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router

}
