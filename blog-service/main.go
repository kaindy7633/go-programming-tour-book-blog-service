package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kaindy7633/go-programming-tour-book/blog-service/global"
	"github.com/kaindy7633/go-programming-tour-book/blog-service/internal/model"
	"github.com/kaindy7633/go-programming-tour-book/blog-service/internal/routers"
	"github.com/kaindy7633/go-programming-tour-book/blog-service/pkg/logger"
	"github.com/kaindy7633/go-programming-tour-book/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	// 自动生成swagger文档 json
	if swagInitErr := exec.Command("swag", "init").Run(); swagInitErr != nil {
		panic(swagInitErr)
	}
	fmt.Println("swagger.json生成成功")
}

// @title 博客系统 (这里是标题)
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目 (描述信息)
// @termsOfService https://github.com/kaindy7633/go-programming-tour-book-blog-service

// @contact.name 联系人信息
// @contact.url https://www.swagger.io/support 网址
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://127.0.0.1:8000
// @BasePath /api/v1
func main() {
	global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		fmt.Println("happened some err: ", err)
	}
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
