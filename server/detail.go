package server

import (
	"github.com/Maymomo/Switch-Harmony/task"
	"github.com/gin-gonic/gin"
	"net/http"
)

type QueryArgs struct {
	ID int `uri:"id" binding:"required"`
}

func GetDetailInfo(r *gin.Context) {
	var args QueryArgs
	if err := r.ShouldBindUri(&args); err != nil {
		r.AbortWithStatus(http.StatusNotFound)
		return
	}
	detail, err := task.GetGameDetailInfo("http://www.eshop-switch.com/game/2.html")
	if err != nil {
		r.AbortWithStatus(http.StatusNotFound)
		return
	}
	r.JSON(http.StatusOK, detail)
}

func GetSummaryInfo(r *gin.Context) {

}
