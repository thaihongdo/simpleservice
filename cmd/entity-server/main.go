package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"simpleservice/cmd/entity-server/model"
	"simpleservice/cmd/entity-server/router"
	"simpleservice/configs"
)

func main() {
	wd, _ := os.Getwd()
	configPath := filepath.Join(wd, "configs/config.toml")
	cf := configs.InitFromFile(configPath)
	// init logger

	// init database connection
	dbPath := filepath.Join(wd, cf.DBConnection)
	closeFunc, _ := model.InitFromSQLLite(dbPath)

	// init router
	routersInit := router.InitRouter(cf.EnvPrefix)

	// init enpoint
	endPoint := fmt.Sprintf(":%d", cf.ServerPort)
	readTimeout := time.Minute
	writeTimeout := time.Minute
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	_ = server.ListenAndServe()
	defer closeFunc()
}
