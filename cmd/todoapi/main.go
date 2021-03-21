package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"todo-api/internal"
	"todo-api/internal/database"
	"todo-api/internal/item"
	"todo-api/internal/transport"
	"todo-api/pkg/config"
)

func migrateDbSchema() {
	log.Info("Migrating database schema")
	dh := database.GetHandler()
	if err := dh.DB().AutoMigrate(&item.Item{}); err != nil {
		panic(`error migrating db schema: ` + err.Error())
	}
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	var (
		itemsManager = item.NewItemsManager(item.NewItemsRepository())
		service      = internal.NewTodoApiService(itemsManager)
		endpoints    = transport.MakeEndpoints(service)
		port         = ":" + strconv.Itoa(config.Conf.Server.Port)
		httpHandler  = transport.NewHTTPHandler(endpoints)
	)

	migrateDbSchema()
	server := &http.Server{
		Addr:    port,
		Handler: httpHandler,
	}
	log.WithFields(log.Fields{"port": port}).Info("starting up service")
	server.ListenAndServe()
}
