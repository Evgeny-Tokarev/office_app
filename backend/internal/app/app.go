package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/evgeny-tokarev/office_app/backend/internal/bootstrap"
	"github.com/evgeny-tokarev/office_app/backend/internal/config"
	"github.com/evgeny-tokarev/office_app/backend/internal/middlware"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories/employee_repository"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories/office_repository"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories/user_repository"
	"github.com/evgeny-tokarev/office_app/backend/internal/services/employeeservice"
	"github.com/evgeny-tokarev/office_app/backend/internal/services/officeservice"
	"github.com/evgeny-tokarev/office_app/backend/internal/services/userservice"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	config config.Config
	store  Store
}

type Store struct {
	EmployeeRepo *employee_repository.Queries
	OfficeRepo   *office_repository.Queries
	UserRepo     *user_repository.Queries
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		EmployeeRepo: employee_repository.New(db),
		OfficeRepo:   office_repository.New(db),
		UserRepo:     user_repository.New(db),
	}
}

func NewApp(config config.Config, tokenType string) (*App, error) {

	db, err := bootstrap.InitSqlDB(config)
	if err != nil {
		return nil, err
	}
	storage := NewStore(db)
	app := &App{
		config: config,
		store:  *storage,
	}

	return app, nil
}

func (a *App) Run(cfg config.Config) error {
	log.Infof("Starting server on address %s and port %d", cfg.PgHost, cfg.Port)
	router := mux.NewRouter()
	authRoutes := router.PathPrefix("/").Subrouter()
	emplService := employeeservice.New(a.store.EmployeeRepo)
	officeService := officeservice.New(a.store.OfficeRepo)
	userService, err := userservice.New(a.store.UserRepo, cfg)
	if err != nil {
		return err
	}

	emplService.SetHandlers(router, authRoutes)
	officeService.SetHandlers(router, authRoutes)
	userService.SetHandlers(router, authRoutes)

	authRoutes.Use(middlware.TokenMiddleware(userService))

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
