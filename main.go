package main

import (
	"flag"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/go-steven/empty-rest-project/handler"
	"github.com/go-steven/empty-rest-project/router"
	"github.com/go-steven/utils"
	log "github.com/kdar/factorlog"
	"math/rand"
	"runtime"
	"time"
)

var (
	logFlag  = flag.String("log", "", "set log path")
	portFlag = flag.Int("port", 10008, "set port")

	logger *log.FactorLog
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	logger = utils.SetGlobalLogger(*logFlag)
	handler.SetLogger(logger)

	gin.SetMode(gin.ReleaseMode)
	r := router.NewRouter()
	logger.Infof("Rest API Server started at:0.0.0.0:%d", *portFlag)
	defer func() {
		logger.Infof("Rest API Server exit from:0.0.0.0:%d", *portFlag)
	}()
	endless.ListenAndServe(fmt.Sprintf(":%d", *portFlag), r)
}
