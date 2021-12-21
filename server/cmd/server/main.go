package main

import (
	"context"
	"fmt"
	"github.com/nickname038/architecture-3/db"
	"github.com/nickname038/architecture-3/menu"
	"github.com/nickname038/architecture-3/orders"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	config := db.DatabaseConnectionConfig{
		DbName:   "postgres",
		User:     "postgres",
		Password: "1234",
		Host:     "127.0.0.1",
		PORT:     5432,
	}
	connection, _ := db.OpenConnection(config)
	menuFacade := menu.NewMenuFacade(connection)
	orderFacade := orders.NewOrderFacade(connection, menuFacade, 0.05, 0.1)

	server := Server{}
	go func() {
		log.Println("Starting chat server...")
		err := server.Start(menuFacade, orderFacade)
		if err == http.ErrServerClosed {
			log.Printf("HTTP server stopped")
		} else {
			log.Fatalf("Cannot start HTTP server: %s", err)
		}
	}()

	// Wait for Ctrl-C signal.
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	<-sigChannel

	if err := server.Stop(); err != nil && err != http.ErrServerClosed {
		log.Printf("Error stopping the server: %s", err)
	}
}

type Server struct {
	server *http.Server
}

func (s *Server) Start(menuFacade *menu.MenuFacade, orderFacade *orders.OrderFacade) error {
	handler := new(http.ServeMux)
	handler.HandleFunc("/menu", menu.HttpHandler(menuFacade))
	handler.HandleFunc("/orders", orders.HttpHandler(orderFacade))
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: handler,
	}
	s.server = server
	return server.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.server.Shutdown(context.Background())
}
