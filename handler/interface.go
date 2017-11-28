package handler

import (
	log "github.com/kdar/factorlog"
)

var (
	Logger *log.FactorLog
)

func SetLogger(l *log.FactorLog) {
	Logger = l
}
