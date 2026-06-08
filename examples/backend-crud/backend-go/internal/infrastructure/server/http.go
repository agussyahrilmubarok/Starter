package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"agussyahrilmubarok.github.io/backend/internal/infrastructure/config"
	"agussyahrilmubarok.github.io/backend/pkg/logger"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	cfg    *config.Config
	router *gin.Engine
	server *http.Server
}

func NewHTTPServer(cfg *config.Config) *HTTPServer {
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.App.Port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	return &HTTPServer{
		cfg:    cfg,
		router: router,
		server: srv,
	}
}

func (s *HTTPServer) Router() *gin.Engine {
	return s.router
}

func (s *HTTPServer) Run() error {
	log := logger.Get()

	go func() {
		log.Info("server running...", zap.String("app_name", s.cfg.App.Name), zap.String("app_addr", s.server.Addr))
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("shutting down server....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	log.Info("server exited gracefully")
	return nil
}
