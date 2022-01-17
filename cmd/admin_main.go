package main

import (
	"gomall/config"
	"gomall/internal/router"
	"gomall/pkg/global/log"
	"net/http"
	"time"
)

func main() {
	// 初始化日志
	log.InitLogger(config.GetConfig().Log.Path, config.GetConfig().Log.Level)
	log.Logger.Info("config", log.Any("config", config.GetConfig()))

	log.Logger.Info("start admin server", log.String("start", "start admin server..."))

	newAdminRouter := router.NewAdminRouter()

	s := &http.Server{
		Addr:           ":8889",
		Handler:        newAdminRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Logger.Error("server error", log.Any("serverError", err))
	}
}
