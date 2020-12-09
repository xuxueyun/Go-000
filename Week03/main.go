package main

/**
 * File :   main.go
 * Author:  xuxueyun
 * Version: 1.0.0
 * Date:    2020/12/9 22:30
 * Copyright: 2020 DanielXU<i@xuxueyun.com>
 * Description:
 */

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/glog"
	"golang.org/x/sync/errgroup"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server!")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 启动 HttpServer
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			glog.Errorf("server listen: %s\n", err)
		}
	}()

	stop := make(chan struct{})
	g, ctx := errgroup.WithContext(context.Background())
	// 当任何 goroutine 产生 error 时，立即关闭 HttpServer
	g.Go(func() error {
		go func() {
			<-ctx.Done()
			ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 延迟 5s 退出
			defer cancel()

			if err := srv.Shutdown(ctx2); err != nil {
				glog.Error("shutdown:", err)
			}
			stop <- struct{}{}
			glog.Warning("HttpServer exit.")
		}()
		return nil
	})

	// 监听系统信号，当接收到退出相关信号退出
	g.Go(func() error {
		quitSig := make(chan os.Signal)
		// 监听到指定信号就给quit
		signal.Notify(quitSig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-quitSig:
				return errors.New("收到退出信号") // 手动退出
			}
		}
	})

	// 后台任务
	g.Go(func() error {
		for {
			<-ctx.Done()
			glog.Warning("ctx Done")
			return ctx.Err()
		}
	})
	err := g.Wait()
	glog.Error(err)
	<-stop
	glog.Info("All Exit.")
}
