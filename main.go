package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"vote/routers"
	"vote/utils/logging"
	"vote/utils/setting"
)

func init() {
	setting.Setup()
	//models.Setup()
	logging.Setup()
}

// @title 投票系统
// @version 1.0
// @description 业务维护资金使用投票系统
// @termsOfService http://123.206.34.236:8000/
// @license.name MIT
// @license.url https://github.com/JiBadBoy/vote
func main() {

	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

}
