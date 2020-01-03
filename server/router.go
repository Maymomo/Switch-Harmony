package server

import "github.com/gin-gonic/gin"

func SetRouter(r *gin.Engine) {
	game := r.Group("/info")
	game.GET("/detail/:id", GetDetailInfo)
	game.GET("/summary/:id", GetSummaryInfo)

}
