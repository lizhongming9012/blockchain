package main

import (
	"NULL/blockchain/models"
	"NULL/blockchain/pkg/logging"
	"NULL/blockchain/pkg/setting"
	"NULL/blockchain/pkg/util"
	"NULL/blockchain/routers"
	v1 "NULL/blockchain/routers/api/v1"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func init() {
	setting.Setup()
	logging.Setup()
	models.Setup()
	util.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dir)
	log.Printf("[info] start http server listening %s", endPoint)

	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		fmt.Println("监听进程启动...")
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("退出", s)
				ExitFunc()
			default:
				fmt.Println("other signal", s)
			}
		}
	}()

	go func() {
		v1.Blockchain, err = models.ReadBlocks()
		if err != nil {
			log.Println("读取[]Block err:", err)
		}
	}()

	err = server.ListenAndServe()
	if err != nil {
		log.Printf("init listen server fail:%v", err)
	}
}

func ExitFunc() {
	fmt.Println("进程断开,开始存储[]Block...")
	if err := models.WriteBlocks(v1.Blockchain); err != nil {
		log.Println("读取[]Block err:", err)
	}
	fmt.Println("存储[]Block完成...")
	os.Exit(0)
}
