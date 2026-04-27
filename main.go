// Package main Product API
//
// Documentation for product API.
//
// Schemes: http
// Host: localhost:9090
// BasePath: /products
// Version: 1.0.0
// License: MIT http://opensource.org/licenses/MIT
// Contact: Jim Rox <jimrox@example.com>
// 
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jimroxodezi/build-ms/handlers"
	"github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)


func main() {
	logger := log.New(os.Stdout, "Server: ", log.LstdFlags)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	productRouter := chi.NewRouter()

	ph := handlers.NewProducts(logger)
	
	productRouter.Get("/products", ph.GetProducts)
	productRouter.Post("/products", ph.AddProduct)
	productRouter.Put("/products/{id}", ph.UpdateProducts)

	r.Mount("/", productRouter)	

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}() // run the server in a goroutine so that it doesn't block the main thread

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	sig := <-stop //block until we receive a signal
	logger.Printf("Received signal: %s", sig)

	logger.Println("Shutting down server...")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancelFunc()

	server.Shutdown(ctx)
	logger.Println("Server gracefully stopped")
}
