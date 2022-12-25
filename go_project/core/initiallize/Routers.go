package initiallize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"louis/global"
	"net/http"
	"strings"
)

type Routers struct {
}

func (r *Routers) DoInitRouters() *gin.Engine {
	Router := gin.Default()

	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOGGER.Info("register swagger handler")

	//todo 大括号换行
	PublicGroup := Router.Group("")
	{
		PublicGroup.GET("/health", func(context *gin.Context) {
			context.JSON(http.StatusOK, "Ok")
		})
	}
	var addr string
	if global.LOUIS_CONFIG.System.Addr != 0 {
		addr = fmt.Sprintf(":%d", global.LOUIS_CONFIG.System.Addr)
	} else {
		s1 := ":"
		s2 := "8080"
		var str []string = []string{s1, s2}
		addr = strings.Join(str, "")

	}
	global.LOGGER.Info("server addr is " + addr)
	r.initServer(addr, Router)
	return Router

}

func (r *Routers) initServer(address string, router *gin.Engine) {
	err := router.Run(address)
	if err != nil {
		panic(err)
	}

}
