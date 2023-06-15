package middlewares

import (
	"final-project-4/database"
	"final-project-4/models"
	"final-project-4/repositories/user_repository/user_pg"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// func ProductAuthorization() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		db := database.GetPostgresInstance()
// 		userData := c.MustGet("userData").(jwt.MapClaims)
// 		userId := uint(userData["id"].(float64))
// 		initialUser := &models.User{}
// 		initialUser.ID = userId

// 		userRepo := user_pg.NewUserPG(db)
// 		err := userRepo.GetUserByID(initialUser)
// 		// abis di Get, objek initialUser akan terupdate,
// 		// smua attribute nya akan terisi.

// 		// user nya fix ada karna udh di cek di authentication,
// 		// tp cek dulu role nya "admin" bukan?
// 		if initialUser.Role != "admin" {
// 			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 				"error":   "Unauthorized",
// 				"message": "You are not allowed to access this product feature",
// 			})
// 			return
// 		}

// 		c.Next()
// 	}
// }

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetPostgresInstance()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		initialUser := &models.User{}
		initialUser.ID = userId

		userRepo := user_pg.NewUserPG(db)
		user, err := userRepo.GetUserByID(initialUser.ID) // Menggunakan initialUser.ID sebagai argumen

		if err != nil {
			// Handle error
			return
		}

		// user nya fix ada karena sudah dicek di authentication,
		// tapi perlu dicek dulu rolenya "admin" bukan?
		if user.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this product feature",
			})
			return
		}

		c.Next()
	}
}



// func CategoryAuthorization() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		db := database.GetPostgresInstance()
// 		userData := c.MustGet("userData").(jwt.MapClaims)
// 		userId := uint(userData["id"].(float64))
// 		initialUser := &models.User{}
// 		initialUser.ID = userId

// 		userRepo := user_pg.NewUserPG(db)
// 		userRepo.GetUserByID(initialUser)
// 		// abis di Get, objek initialUser akan terupdate,
// 		// smua attribute nya akan terisi.

// 		// user nya fix ada karna udh di cek di authentication,
// 		// tp cek dulu role nya "admin" bukan?
// 		if initialUser.Role != "admin" {
// 			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 				"error":   "Unauthorized",
// 				"message": "You are not allowed to access this category feature",
// 			})
// 			return
// 		}

// 		c.Next()
// 	}
// }

func CategoryAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetPostgresInstance()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		initialUser := &models.User{}
		initialUser.ID = userId

		userRepo := user_pg.NewUserPG(db)
		user, err := userRepo.GetUserByID(initialUser.ID) // Menggunakan initialUser.ID sebagai argumen

		if err != nil {
			// Handle error
			return
		}

		// user nya fix ada karena sudah dicek di authentication,
		// tapi perlu dicek dulu rolenya "admin" bukan?
		if user.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this category feature",
			})
			return
		}

		c.Next()
	}
}


func TaskAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetPostgresInstance()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		initialUser := &models.User{}
		initialUser.ID = userId

		userRepo := user_pg.NewUserPG(db)
		user, err := userRepo.GetUserByID(initialUser.ID) // Menggunakan initialUser.ID sebagai argumen
		if err != nil {
			// Handle error
			return
		}

		// user nya fix ada karena sudah dicek di authentication,
		// tapi perlu dicek dulu rolenya "admin" bukan?
		if user.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this task feature",
			})
			return
		}

		c.Next()
	}
}

