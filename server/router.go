package server

import "github.com/gin-gonic/gin"

func SetRouter(r *gin.Engine) {
	game := r.Group("/game")
	game.GET("/:id/detail", GetDetailInfo)
	game.GET("/:id/summary", GetSummaryInfo)
	game.GET("/", GetSummaryOverPage)
}
