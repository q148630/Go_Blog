package main

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/grace"
	"go_blog/pkg/setting"
	"go_blog/routers"
	"log"
	"os"
	"os/signal"
	"time"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in query
// @name token
func main() {

	grace.DefaultReadTimeOut = setting.ReadTimeout
	grace.DefaultWriteTimeOut = setting.WriteTimeout
	grace.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := grace.NewServer(endPoint, routers.InitRouter())

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Server err: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")


}
