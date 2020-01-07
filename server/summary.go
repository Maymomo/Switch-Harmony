package server

import (
	"github.com/Maymomo/Switch-Harmony/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PagingArgs struct {
	Limit  int `form:"limit" binding:"required"`
	Offset int `form:"offset"`
	Query  string `form:"q"`
}
type PagingResp struct {
	List  []*db.GameSummary
	TotalPages int
}

func GetSummaryOverPage(r *gin.Context) {
	var args PagingArgs
	if err := r.ShouldBindQuery(&args); err != nil {
		r.AbortWithStatus(http.StatusNotFound)
		return
	}
	list, count := db.GetSummaryByPage(args.Query, args.Offset, args.Limit)
	r.JSON(http.StatusOK, PagingResp{
		List:  list,
		TotalPages: count,
	})
}

type QuerySummaryArgs struct {
	ID int `uri:"id" binding:"required"`
}

func GetSummaryInfo(r *gin.Context) {
	var args QuerySummaryArgs
	if err := r.ShouldBindUri(&args); err != nil {
		r.AbortWithStatus(http.StatusNotFound)
		return
	}
	ret, err := db.GetSummaryByID(args.ID)
	if err != nil {
		r.AbortWithStatus(http.StatusNotFound)
		return
	}
	r.JSON(http.StatusOK, ret)
}
