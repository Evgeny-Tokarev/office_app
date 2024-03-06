package app

import (
	"context"
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
	"github.com/evgeny-tokarev/office_app/backend/internal/token"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var secret string

type Server struct {
	config     config.Config
	tokenMaker token.Maker
}

func NewServer(config config.Config, tokenType string) (*Server, error) {
	tokenMaker, err := token.NewMaker(tokenType, config.JwtSecret)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

func (s *Server) Run(cfg config.Config) error {
	log.Info("Config: ", cfg)
	secret = cfg.JwtSecret
	db, err := bootstrap.InitSqlDB(cfg)
	if err != nil {
		return err
	}

	router := mux.NewRouter()
	employeeQueries := employee_repository.New(db)
	officeQueries := office_repository.New(db)
	userQueries := user_repository.New(db)
	emplService := employeeservice.New(*employeeQueries)
	officeService := officeservice.New(officeQueries)
	userService, err := userservice.New(userQueries, cfg)
	if err != nil {
		return err
	}

	emplService.SetHandlers(router)
	officeService.SetHandlers(router)
	userService.SetHandlers(router)

	router.Use(middlware.TokenMiddleware(userService))

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
