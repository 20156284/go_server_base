package router

import (
	"github.com/gogf/gf/net/ghttp"
	"go_base_server/server/app/api/response"
	api "go_base_server/server/app/api/system"
	"go_base_server/server/interfaces"
)

type base struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewBaseGroup(router *ghttp.RouterGroup) interfaces.Router {
	return &base{router: router, response: &response.Handler{}}
}

func (b *base) Init() {
	group := b.router.Group("/base")
	{
		group.POST("captcha", b.response.Handler()(api.Base.Captcha))
		group.POST("login", api.GfJWTMiddleware.LoginHandler) // 登录
	}

	db := b.router.Group("/init")
	{
		db.POST("initdb", b.response.Handler()(api.Base.Captcha))
		db.POST("checkdb", api.GfJWTMiddleware.LoginHandler) // 登录
	}

}
