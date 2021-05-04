package sgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/common-go/cgin"
	"github.com/suzuito/s2-demo-go/resource"
	"github.com/suzuito/s2-demo-go/setting"
	"github.com/suzuito/s2-demo-go/sgcp"
)

func middlewareResource(env *setting.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		gcpResource, err := cgin.GetGCPResource(c)
		if err != nil {
			abortError(c, err)
			return
		}
		defer gcpResource.Close()
		r := resource.Resource{
			ArticleStore: sgcp.NewArticleStore(
				gcpResource.GCS,
				env.GCPBucketArticle,
			),
		}
		setCtxResource(c, &r)
		c.Next()
	}
}
