package web

import (
	"github.com/gin-gonic/gin"
	"MBlog/util"
	"MBlog/web/router"
)

var engine *gin.Engine

func init() {
	if util.BlogConfig.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine = gin.New()
	engine.Use(gin.Recovery())

	engine.SetFuncMap(util.GetFuncMap())
	engine.LoadHTMLFiles(util.GetTemplateFiles()...)

	router.RegisterRoute(engine)
}

func Run() {
	engine.Run(util.BlogConfig.Web.Listen)
}