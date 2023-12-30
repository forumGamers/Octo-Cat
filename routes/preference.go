package routes

import (
	"github.com/forumGamers/Octo-Cat/controllers"
	md "github.com/forumGamers/Octo-Cat/middlewares"
	"github.com/gin-gonic/gin"
)

func (r *routes) preferenceRoute(rg *gin.RouterGroup, mds md.Middleware, preference controllers.PreferenceController) {
	rg.Group("/preference").
		Use(mds.Authentication).
		POST("/", preference.CreateData)
}
