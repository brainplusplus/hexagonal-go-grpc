package echo

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
	"time"
)

// Server serves HTTP endpoints.
type Server interface {
	Run() chan error
	Router() *echo.Echo
	Close() error
}

type server struct {
	echo   *echo.Echo
	server *http.Server
	cfg    Config
}

// Config is basic HTTP server config.
type Config struct {
	Port            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	GracefulTimeout time.Duration
}

// New to create a new Echo web server.
func NewServer(cfg Config) Server {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &server{
		echo: e,
		cfg:  cfg,
	}
}

// Router returns the Echo router.
func (s *server) Router() *echo.Echo {
	return s.echo
}

// Run starts the Echo server.
func (s *server) Run() chan error {
	ch := make(chan error)
	go s.run(ch)
	return ch
}

func (s *server) run(ch chan error) {
	h2s := &http2.Server{}
	s.server = &http.Server{
		Addr:         ":" + s.cfg.Port,
		Handler:      h2c.NewHandler(s.echo, h2s),
		ReadTimeout:  s.cfg.ReadTimeout,
		WriteTimeout: s.cfg.WriteTimeout,
	}

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		ch <- err
	}
}

// Close stops the server gracefully.
func (s *server) Close() error {
	if s.server == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.GracefulTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
