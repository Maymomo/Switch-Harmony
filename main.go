package main

import (
	"github.com/Maymomo/Switch-Harmony/server"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	server.SetRouter(r)
	r.Run()
}
