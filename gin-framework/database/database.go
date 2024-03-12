package database

import (
	"errors"
	"strconv"

	"github.com/Emrul-Hasan-Emon/demo-go-gin/model"
)

func GetArticleList() []model.Article {
	return getAllArtilces()
}

func getAllArtilces() []model.Article {
	var articleList = []model.Article{
		model.Article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
		model.Article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
		model.Article{ID: 3, Title: "Article 3", Content: "Article 3 body"},
	}
	return articleList
}
func GetArticleById(id string) (*model.Article, error) {
	articleList := getAllArtilces()
	for _, article := range articleList {
		if strconv.Itoa(article.ID) == id {
			return &article, nil
		}
	}
	return nil, errors.New("Article not found")
}
