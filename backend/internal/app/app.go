package app

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"lesson4/internal/bootstrap"
	"lesson4/internal/config"
	"lesson4/internal/repositories/employeerepository/employeesql"
	"lesson4/internal/repositories/officerepository/officessql"
	"lesson4/internal/services/employeeservice"
	"lesson4/internal/services/officeservice"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg config.Config) error {

	db, err := bootstrap.InitSqlDB(cfg)
	if err != nil {
		return err
	}

	router := mux.NewRouter()
	emplService := employeeservice.New(employeesql.New(db))
	officeService := officeservice.New(officessql.New(db))

	emplService.SetHandlers(router)
	officeService.SetHandlers(router)

	// CORS
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	method := handlers.AllowedMethods([]string{"POST", "GET", "DELETE", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})

	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		Handler: handlers.CORS(header, method, origins)(router),
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	gracefullyShutdown(ctx, cancel, server)
	return nil
}

func gracefullyShutdown(ctx context.Context, cancel context.CancelFunc, server *http.Server) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)
	<-ch
	if err := server.Shutdown(ctx); err != nil {
		log.Warning(err)
	}
	cancel()
}
