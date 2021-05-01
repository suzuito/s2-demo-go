package sgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/common-go/cgin"
	"github.com/suzuito/s2-demo-go/resource"
)

const ctxResource = "resource"

func setCtxResource(c *gin.Context, v *resource.Resource) {
	c.Set(ctxResource, v)
}

func getCtxResource(c *gin.Context) (*resource.Resource, error) {
	r, err := cgin.GetCtxVariable(c, ctxResource, &resource.Resource{})
	if err != nil {
		return nil, err
	}
	return r.(*resource.Resource), nil
}
