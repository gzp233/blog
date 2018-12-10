package route

import (
	"blog/control"
	"blog/model"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Run 入口
func Run() {
	engine := echo.New()
	engine.Renderer = initRender()
	engine.Use(middleware.Recover())
	engine.Use(middleware.LoggerWithConfig(logConfig))
	engine.Use(middleware.CORSWithConfig(crosConfig)) //允许跨域
	engine.HideBanner = true
	engine.HTTPErrorHandler = HTTPErrorHandler
	engine.Static(`/res`, "res")
	// 前后端分离页面
	engine.File(`/core/*`, "res/dist/index.html")

	engine.GET(`/`, control.IndexView)
	engine.GET(`/archives`, control.ArchivesView)
	engine.GET(`/tags`, control.TagsView)
	engine.GET(`/tag/:tag`, control.TagPostView)
	engine.GET(`/cate/:cate`, control.CatePostView)
	engine.GET(`/about`, control.AboutView)
	engine.GET(`/links`, control.LinksView)
	engine.GET(`/post/*`, control.PostView)
	engine.GET(`/page/*`, control.PageView)
	engine.File(`/favicon.ico`, "favicon.ico")

	engine.GET(`/user/exist/:num`, control.UserExist)
	engine.POST(`/login`, control.UserLogin)
	engine.POST(`/logout`, control.UserLogout)

	// 需要登陆才能访问
	app := engine.Group(``, midJwt)
	app.GET(`/auth`, control.UserAuth)
	app.POST(`/upload`, control.Upload)
	app.POST(`/user/edit/self`, control.UserEditSelf)
	app.POST(`/user/pass`, control.UserPass)
	adm := app.Group(`/adm`)
	adm.GET(`/cate/all`, control.CateAll)
	adm.GET(`/cate/del/:id`, control.CateDel)
	adm.GET(`/cate/post/:cid`, control.CatePost)
	adm.POST(`/cate/add`, control.CateAdd)
	adm.POST(`/cate/edit`, control.CateEdit)
	// 文章
	adm.GET(`/post/get/:id`, control.PostGet)
	adm.GET(`/post/tag/ids/:id`, control.PostTagIds)
	adm.POST(`/post/opts`, control.PostOpts)
	adm.GET(`/post/del/:id`, control.PostDel)
	//页面
	adm.GET(`/page/all`, control.PostPageAll)
	//标签
	adm.GET(`/tag/all`, control.TagAll)
	adm.GET(`/tag/del/:id`, control.TagDel)
	adm.POST(`/tag/add`, control.TagAdd)
	adm.POST(`/tag/edit`, control.TagEdit)
	// 配置
	adm.GET(`/opts/base`, control.OptsBase)
	adm.GET(`/opts/:key`, control.OptsGet)
	adm.POST(`/opts/edit`, control.OptsEdit)
	engine.Start(model.Conf.Addr)
}
