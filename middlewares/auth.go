package middlewares

import (
	"log"
	"net/http"

	"example.com/events/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		log.Println("here 1")
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized action"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		log.Println("here 1")
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized action"})
		return
	}

	context.Set("userId", userId)

	context.Next()

}
