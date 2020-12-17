package web

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Router       *gin.Engine
}

func NewServer(port string, readTimeout time.Duration, writeTimeout time.Duration, router *gin.Engine) *Server {
	return &Server{
		Port:         port,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		Router: router,
	}
}

func (s *Server) Start() {
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", s.Port),
		Handler:      s.Router,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	configGracefulShutdown(httpServer)
}

func configGracefulShutdown(httpServer *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
