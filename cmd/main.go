package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"net/http"

	"todo-api/internal"
	"todo-api/internal/transport"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	var (
		service     = internal.NewTodoApiService()
		endpoints   = transport.MakeEndpoints(service)
		port        = flag.String("port", ":8080", "port to run the service on")
		httpHandler = transport.NewHTTPHandler(endpoints)
	)

	flag.Parse()
	server := &http.Server{
		Addr:    *port,
		Handler: httpHandler,
	}
	log.WithFields(log.Fields{"port": *port}).Info("starting up service")
	server.ListenAndServe()
}
