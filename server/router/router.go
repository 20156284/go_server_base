package router

import (
	"github.com/gogf/gf/frame/g"
	extra "go_base_server/server/router/extra"
	"go_base_server/server/router/internal"
	system "go_base_server/server/router/system"
)

var Routers = new(routers)

type routers struct{}

func (r *routers) Init() {
	public := g.Server().Group("")
	{ // 无需鉴权中间件
		system.NewBaseGroup(public).Init()
		system.NewConfigRouter(public).Init()
	}
	private := g.Server().Group("").Middleware(internal.Middleware.JwtAuth, internal.Middleware.CasbinRbac)
	{ // 需要Jwt鉴权, casbin鉴权
		system.NewApiRouter(private).Init()
		system.NewAdminRouter(private).Init()
		system.NewMenuRouter(private).Init()
		system.NewEmailRouter(private).Init()
		system.NewCasbinRouter(private).Init()
		system.NewGenerateRouter(private).Init()
		system.NewAuthorityRouter(private).Init()
		system.NewDictionaryRouter(private).Init()
		system.NewJwtBlacklistRouter(private).Init()
		system.NewOperationRecordRouter(private).Init()
		system.NewDictionaryDetailRouter(private).Init()

		extra.NewFileRouter(private).Init()
		extra.NewExcelRouter(private).Init()
		extra.NewWorkflowRouter(private).Init()
		extra.NewSimpleUploaderRouter(private).Init()
	}
}
