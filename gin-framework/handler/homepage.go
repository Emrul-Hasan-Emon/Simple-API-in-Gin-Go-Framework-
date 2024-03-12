package handler

import (
	"github.com/Emrul-Hasan-Emon/demo-go-gin/database"
	"github.com/Emrul-Hasan-Emon/demo-go-gin/render"
	"github.com/gin-gonic/gin"
)

func ShowHomePage(ctx *gin.Context) {
	articleList := database.GetArticleList()
	// ctx.HTML(
	// 	http.StatusOK,
	// 	"index.html",
	// 	gin.H{
	// 		"title":   "Home Page",
	// 		"payload": articleList,
	// 	},
	// )
	render.Render(ctx, gin.H{
		"title":   "Home Page",
		"payload": articleList,
	}, "index.html")
}
