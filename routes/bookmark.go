package routes

import (
	"github.com/forumGamers/Octo-Cat/controllers"
	md "github.com/forumGamers/Octo-Cat/middlewares"
	"github.com/gin-gonic/gin"
)

func (r routes) bookmarkRoutes(rg *gin.RouterGroup, mds md.Middleware, bookmark controllers.BookmarkController) {
	uri := rg.Group("/bookmark")

	uri.Use(mds.SetContexts)
	uri.Use(mds.Authentication)
	uri.POST("/:postId", bookmark.CreateBookmark)
}
