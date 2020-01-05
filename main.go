package main

import (
	"fmt"
	"github.com/Maymomo/Switch-Harmony/common"
	"github.com/Maymomo/Switch-Harmony/db"
	"github.com/Maymomo/Switch-Harmony/server"
	"github.com/Maymomo/Switch-Harmony/task"
	"github.com/gin-gonic/gin"
)

func main() {
	config := common.GetConfig().WebConfig
	db.DataBaseInit()
	r := gin.Default()
	server.SetRouter(r)
	addr := fmt.Sprintf("%s:%d", config.Address, config.Port)
	task.BackgroundWork()
	r.Run(addr)
}
