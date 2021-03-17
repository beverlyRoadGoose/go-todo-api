package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"net/http"
	"todo-api/internal/database"
	"todo-api/internal/item"

	"todo-api/internal"
	"todo-api/internal/transport"
)

func migrateDbSchema() {
	log.Info("Migrating database schema")
	dh := database.GetHandler()
	if err := dh.DB().AutoMigrate(&item.Item{}); err != nil {
		panic(`error migrating db schema:` + err.Error())
	}
}

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
	migrateDbSchema()
	server := &http.Server{
		Addr:    *port,
		Handler: httpHandler,
	}
	log.WithFields(log.Fields{"port": *port}).Info("starting up service")
	server.ListenAndServe()
}
