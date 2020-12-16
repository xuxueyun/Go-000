package myapp

import (
	"Week04/global"
	"Week04/pkg/config"
	"Week04/pkg/database"
	"Week04/pkg/utils"
	"Week04/router"

	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gogf/gf/os/glog"
	"github.com/spf13/cobra"
)

/**
 * File :   main.go
 * Author:  xuxueyun
 * Version: 1.0.0
 * Date:    2020/12/16 19:24
 * Copyright: 2020 DanielXU<i@xuxueyun.com>
 * Description:
 */

var (
	configYml string
	port      string
	mode      string
	StartCmd  = &cobra.Command{
		Use:          "myapp",
		Short:        "start myapp server",
		Example:      "myapp -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "HTTP Server listening on")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg: dev,test,prod")
}

func setup() {
	// 1. 读取配置
	config.Setup(configYml)

	// 2. 初始化数据库链接
	database.Setup(config.DatabaseConfig.Driver)
}

func run() error {
	r := router.InitRouter()
	defer global.Eloquent.Close()

	srv := &http.Server{
		Addr:    config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			glog.Fatal("listen:", err)
		}
	}()
	fmt.Println(utils.Green("Server run at:"))
	fmt.Printf("-  Local:   http://localhost:%s/ \r\n", config.ApplicationConfig.Port)
	fmt.Printf("-  Network: http://%s:%s/ \r\n", utils.GetLocaHonst(), config.ApplicationConfig.Port)

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP, syscall.SIGQUIT)
	sig := <-sigs
	fmt.Printf("%s Shutdown Server ... \r\n", utils.Red(sig.String()))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		glog.Fatal("Server Shutdown:", err)
	}
	glog.Println("Server exiting")

	return nil
}

func tip() {

	glog.Info("Version:   ", utils.Green(global.Version))
	glog.Info("BuildTime: ", utils.Green(global.BuildTime))
	glog.Info("GitHash:   ", utils.Green(global.GitHash))
	glog.Info("go version:", utils.Green(global.GoVersion))
}
