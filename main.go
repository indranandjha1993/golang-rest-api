package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	. "site-api/config"
	"site-api/routers"
	"time"
)

var config Config

func init() {
	config.Read()
}

func main() {
	log.SetOutput(os.Stdout)
	logFormatter := new(LogFormatter)
	logFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logFormatter.LevelDesc = []string{"PANIC", "FATAL", "ERROR", "WARN", "INFO", "DEBUG", "TRACE"}
	log.SetFormatter(logFormatter)

	fmt.Printf("Site configuration: %v", config)
	r := routers.Routers()
	srv := &http.Server{
		Handler:      r,
		Addr:         config.Server.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
