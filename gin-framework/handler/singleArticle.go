package handler

import (
	"net/http"

	"github.com/Emrul-Hasan-Emon/demo-go-gin/database"
	"github.com/Emrul-Hasan-Emon/demo-go-gin/render"
	"github.com/gin-gonic/gin"
)

func GetSingleArtilce(ctx *gin.Context) {
	articleID := ctx.Param("article_id")
	article, err := database.GetArticleById(articleID)

	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}

	// ctg.HTML(
	// 	http.StatusOK,
	// 	"article.html",

	// 	gin.H{
	// 		"title":   article.Title,
	// 		"payload": article,
	// 	},
	// )
	render.Render(ctx, gin.H{
		"title":   article.Title,
		"payload": article,
	}, "article.html")
}
