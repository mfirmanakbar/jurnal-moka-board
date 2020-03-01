package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mfirmanakbar/jurnal-moka-board/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	port := ":8282"
	routes()
	logger.Info(fmt.Sprintf("Started Jurnal Moka Board on %s ...", port))
	router.Run(port)
}
