package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserID(ctx *gin.Context) (string, error) {
	user, exists := ctx.Get("user")
	if !exists {
		return "", errors.New("user not found in context")
	}

	userID, ok := user.(string)
	if !ok {
		return "", errors.New("user ID is not a string")
	}

	return userID, nil
}
