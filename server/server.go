package server

import (
	"github.com/gin-gonic/gin"

	bihin "Goods-Management-System-API/controller"
)

// サーバー起動
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/bihins")
	{
		ctrl := bihin.Controller{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	return r
}
