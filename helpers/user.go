package helper

import (
	"github.com/forumGamers/Octo-Cat/pkg/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GetUser(c *gin.Context) user.User {
	var user user.User

	claimMap, ok := c.Get("user")
	if !ok {
		return user
	}

	claim, oke := claimMap.(jwt.MapClaims)
	if !oke {
		return user
	}

	for key, val := range claim {
		switch key {
		case "UUID":
			user.UUID = val.(string)
		case "loggedAs":
			user.LoggedAs = val.(string)
		}
	}
	return user
}
