package main

import (
	"github.com/dprapas/go-todo-service/config"
	"github.com/dprapas/go-todo-service/ports"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {

	config.Load()
	l := log.New()
	l.SetFormatter(&log.JSONFormatter{})
	l.SetLevel(log.InfoLevel)

	srv := &http.Server{
		Handler: ports.InitRouter(l),
		Addr:    ":" + config.Configuration.Server.Port,
		WriteTimeout: time.Duration(config.Configuration.Server.WriteTimeoutMS) * time.Millisecond,
		ReadTimeout: time.Duration(config.Configuration.Server.ReadTimeoutMS) * time.Millisecond,
	}
	l.Info("Todo Service Started...")
	l.Fatal(srv.ListenAndServe())
}
