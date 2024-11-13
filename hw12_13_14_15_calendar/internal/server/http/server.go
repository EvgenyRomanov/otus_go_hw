package internalhttp

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	host   string
	port   int
	logger Logger
	app    Application
	server *http.Server
}

type Logger interface {
	Debug(msg string, params ...any)
	Info(msg string, params ...any)
	Warning(msg string, params ...any)
	Error(msg string, params ...any)
}

type Application interface { // TODO
}

func NewServer(host string, port int, logger Logger, app Application) *Server {
	return &Server{
		host:   host,
		port:   port,
		logger: logger,
		app:    app,
	}
}

func (s *Server) Start(ctx context.Context) error {
	// router init
	r := s.initRouter()

	// setup logging middleware
	handlerMiddleware := loggingMiddleware(r, s.logger)

	go func() {
		<-ctx.Done()
		s.Stop(ctx)
	}()

	s.server = &http.Server{
		Addr:              fmt.Sprintf("%s:%d", s.host, s.port),
		Handler:           handlerMiddleware,
		ReadHeaderTimeout: 20 * time.Second,
	}

	s.logger.Info("http-server is up...")

	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) initRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", s.defaultHandler).Methods(http.MethodGet)

	return r
}

func (s *Server) defaultHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
