package main

import (
	"fmt"
	"net/http"
	"time"

	"test-project-hernan/config"
	"test-project-hernan/src"
	"test-project-hernan/src/libs/env"
	"test-project-hernan/src/libs/logger"
)

func main() {
	config.SetupCommonDependencies()
	defer config.TearDownCommonDependencies()
	handler := src.SetupHandler()

	host := fmt.Sprint(":", env.TestProjectServiceRestPort)
	srv := &http.Server{
		Handler:      *handler,
		Addr:         host,
		WriteTimeout: 20 * time.Second,
		ReadTimeout:  20 * time.Second,
	}
	go srv.ListenAndServe()
	logger.GetInstance().Info("Server Listening on ", host)

	select {} //Infinite waiting
}
