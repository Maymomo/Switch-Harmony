package server

import (
	"github.com/Maymomo/Switch-Harmony/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

const detailURL = "http://www.eshop-switch.com/game/"

type QueryArgs struct {
	ID int `uri:"id" binding:"required"`
}

func GetDetailInfo(r *gin.Context) {
	var args QueryArgs
	if err := r.ShouldBindUri(&args); err != nil {
		r.AbortWithStatus(http.StatusNotFound)
		return
	}
	detail, err := db.QueryDetailByID(args.ID)
	if err != nil {
		r.AbortWithStatus(http.StatusNotFound)
		return
	}
	r.JSON(http.StatusOK, detail)
}
