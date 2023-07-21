package initiallize

import (
	"fmt"
	"louis/global"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Routers struct {
}

func DoInitRouters() *gin.Engine {
	Router := gin.Default()

	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logger.Info("register swagger handler")

	//todo 大括号换行
	PublicGroup := Router.Group("")
	{
		PublicGroup.GET("/health", func(context *gin.Context) {
			context.JSON(http.StatusOK, "Ok")
		})
	}
	r := Routers{}

	r.initServer(Router)

	return Router

}

func (r *Routers) initServer(router *gin.Engine) {
	var addr string
	if global.LOUIS_CONFIG.System.Addr != 0 {
		addr = fmt.Sprintf(":%d", global.LOUIS_CONFIG.System.Addr)
	} else {
		s1 := ":"
		s2 := "8080"
		var str []string = []string{s1, s2}
		addr = strings.Join(str, "")
	}
	logger.Info("server addr is " + addr)
	err := router.Run(addr)
	if err != nil {
		panic(err)
	}
	logger.Info("server run ok")
	return
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
