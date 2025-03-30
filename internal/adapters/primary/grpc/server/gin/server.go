package gin

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// Server serves HTTP endpoints.
type Server interface {
	Run() chan error
	Router() *gin.Engine
	Close() error
}

type server struct {
	router *gin.Engine
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

// NewServer creates a new Gin web server.
func NewServer(cfg Config) Server {
	r := gin.Default()

	return &server{
		router: r,
		cfg:    cfg,
	}
}

// Router returns the Gin router.
func (s *server) Router() *gin.Engine {
	return s.router
}

// Run starts the Gin server.
func (s *server) Run() chan error {
	ch := make(chan error)
	go s.run(ch)
	return ch
}

func (s *server) run(ch chan error) {
	h2s := &http2.Server{}
	s.server = &http.Server{
		Addr:         ":" + s.cfg.Port,
		Handler:      h2c.NewHandler(s.router, h2s),
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
