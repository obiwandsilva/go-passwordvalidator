package web

import (
	"github.com/bnkamalesh/webgo/v4"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

func (s *Server) Start(port string, routes []*webgo.Route) {
	router := webgo.NewRouter(&webgo.Config{
		Host: "",
		Port: port,
		ReadTimeout: s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}, routes)

	configGracefulShutdown(router)

	router.Start()
}

func configGracefulShutdown(router *webgo.Router) {
	osSig := make(chan os.Signal, 5)

	go func() {
		<-osSig
		log.Println("shutdown signal received")
		// Initiate HTTP server shutdown
		err := router.Shutdown()
		if err != nil {
			log.Printf("error when shutting down: %v\n", err)
			os.Exit(1)
		} else {
			log.Println("shutdown complete")
			os.Exit(0)
		}
	}()

	signal.Notify(osSig, os.Interrupt, syscall.SIGTERM)
}
