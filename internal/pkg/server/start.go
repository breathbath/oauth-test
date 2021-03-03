package server

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var (
	defaultPort  = ":8000"
	WriteTimeout = 15 * time.Second
	ReadTimeout  = 15 * time.Second
)

func Start(handlers map[string]http.Handler) error {
	router := mux.NewRouter()

	for uri, handl := range handlers {
		router.Handle(uri, handl)
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         defaultPort,
		WriteTimeout: WriteTimeout,
		ReadTimeout:  ReadTimeout,
	}

	logrus.Infof("Started server at %s", defaultPort)
	return srv.ListenAndServe()
}
