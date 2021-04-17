package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// Server is the HTTP server structure
type Server struct {
	logger InfoLogger
	config *Config
	admin  Admin
	engine *gin.Engine
}

// mountAdmin mount the admin
func (s *Server) mountAdmin() {
	s.engine.Any(
		fmt.Sprintf("/%s/*resources", s.admin.GetBasePath()),
		gin.WrapH(s.admin.Mount()),
	)
}

// initLogger setups the logger
func (s *Server) initLogger() {
	file, err := os.OpenFile(
		s.config.ServerLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600,
	)
	if err != nil {
		panic(errors.Wrap(err, "logger"))
	}
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

// getCORS return the CORS middleware
func (s *Server) getCORS() cors.Config {
	config := cors.DefaultConfig()
	if s.config.IsReleaseMode {
		config.AllowOrigins = s.config.ServerAllowOrigins
	} else {
		config.AllowAllOrigins = true
	}
	config.AddAllowHeaders("Authorization")
	return config
}

// setEngine sets the gin engine
func (s *Server) setEngine() {
	s.initLogger()
	if s.config.IsReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	s.engine = gin.Default()
	s.mountAdmin()
	s.engine.Use(cors.New(s.getCORS()))
}

// Run runs the admin
func (s *Server) Run(ctx context.Context, args []string) {
	s.setEngine()
	httpServer := &http.Server{
		Addr: fmt.Sprintf(
			"%s:%s", s.config.ServerHost, s.config.ServerPort,
		),
		Handler:      s.engine,
		ReadTimeout:  s.config.ServerReadTimeout,
		WriteTimeout: s.config.ServerWriteTimeout,
	}

	go func() {
		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(errors.Wrap(err, "server"))
		}
	}()
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case result := <-quit:
		s.logger.Infof("Shutdown the HTTP server. Signal: %v", result)
	case <-ctx.Done():
	}
	shutdownCtx, cancel := context.WithTimeout(
		context.Background(), s.config.ServerShutdownTimeout,
	)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		panic(errors.Wrap(err, "server"))
	}
}

// NewServer returns a new server object
func NewServer(a Admin, l InfoLogger) *Server {
	config := NewConfig()
	s := Server{config: config, admin: a, logger: l}
	return &s
}
