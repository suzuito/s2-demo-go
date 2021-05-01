package sgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/s2-demo-go/entity"
	"github.com/suzuito/s2-demo-go/usecase"
)

func handlerArticles() gin.HandlerFunc {
	return func(c *gin.Context) {
		resource, _ := getCtxResource(c)
		articleList := entity.ArticleListItem{}
		if err := usecase.GetArticleList(c, resource.ArticleStore, &articleList); err != nil {
			abortError(c, err)
			return
		}
		c.JSON(http.StatusOK, usecase.NewResponseArticleListItem(&articleList))
	}
}

func handlerArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		resource, _ := getCtxResource(c)
		articleID := entity.ArticleID(c.Param("articleID"))
		article := entity.Article{}
		if err := usecase.GetArticle(c, resource.ArticleStore, articleID, &article); err != nil {
			abortError(c, err)
			return
		}
		c.JSON(http.StatusOK, usecase.NewResponseArticle(&article))
	}
}
