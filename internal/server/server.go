package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type server struct {
	srv *http.Server
	l   *slog.Logger
}

type ServerOption func(*server)

func WithAddr(addr string) ServerOption {
	return func(s *server) {
		s.srv.Addr = addr
	}
}

func WithTLSConfig(tls *tls.Config) ServerOption {
	return func(s *server) {
		s.srv.TLSConfig = tls
	}
}

func WithReadTimeout(readTimeout time.Duration) ServerOption {
	return func(s *server) {
		s.srv.ReadTimeout = readTimeout
	}
}

func WithWriteTimeout(writeTimeout time.Duration) ServerOption {
	return func(s *server) {
		s.srv.WriteTimeout = writeTimeout
	}
}

func New(handler http.Handler, logger *slog.Logger, opts ...ServerOption) *server {
	server := new(server)
	server.srv = &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	server.l = logger.WithGroup("server")

	for _, opt := range opts {
		opt(server)
	}

	return server
}

func (s *server) Start() error {
	s.l.Info(fmt.Sprintf("Server is starting on %s ...", s.srv.Addr))
	return s.srv.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	s.l.Info("Server is shutting down ...")
	return s.srv.Shutdown(ctx)
}

func (s *server) ServeTLS(certFile, keyFile string) error {
	return s.srv.ListenAndServeTLS(certFile, keyFile)
}

func (s *server) RunWithGracefulShutdown() {
	go func() {
		if err := s.Start(); err != nil && err != http.ErrServerClosed {
			s.l.Error(fmt.Sprintf("Server failed to start :%v", err))
			panic(err)
		}
	}()

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	<-gracefulStop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		s.l.Error(fmt.Sprintf("Server failed to shutdown :%v", err))
		panic(err)
	}
}
