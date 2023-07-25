package middlewares

import (
	"final-project-4/database"
	"final-project-4/repositories/user_repository/user_pg"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetPostgresInstance()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))

		userRepo := user_pg.NewUserPG(db)
		initialUser, err := userRepo.GetUserByID(userId)
		if err != nil {
			// Handle error ketika mendapatkan user
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": "Failed to retrieve user data",
			})
			return
		}

		// abis di Get, objek initialUser akan terupdate,
		// smua attribute nya akan terisi.

		// user nya fix ada karna udh di cek di authentication,
		// tp cek dulu role nya "admin" bukan?
		if initialUser.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this product feature",
			})
			return
		}

		c.Next()
	}
}
