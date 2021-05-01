package sgin

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/common-go/cgcp"
	"github.com/suzuito/common-go/cgin"
	"github.com/suzuito/s2-demo-go/setting"
)

func NewRoute(r *gin.Engine, env *setting.Env, genGCPResource *cgcp.GCPContextResourceGenerator) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: strings.Split(env.AllowedOrigins, ","),
		AllowMethods: strings.Split(env.AllowedMethods, ","),
		AllowHeaders: strings.Split(env.AllowedHeaders, ","),
	}))
	r.Use(cgin.MiddlewareGCPResource(genGCPResource))
	r.Use(middlewareResource(env))

	{
		articles := r.Group("articles")
		articles.GET("", handlerArticles())
		articles.GET(":articleID", handlerArticle())
	}
}
