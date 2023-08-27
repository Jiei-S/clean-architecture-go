package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env"

	rest "github.com/Jiei-S/boilerplate-clean-architecture/go-rest/internal/infrastructure/openapi"

	chi_middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
)

type ServerEnv struct {
	Port int `env:"API_PORT"`
}

func main() {
	cfg := ServerEnv{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	swagger, err := rest.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	swagger.Servers = nil

	logger := httplog.NewLogger("server", httplog.Options{
		JSON: true,
	})

	r := chi.NewRouter()
	c := Init()
	r.Use(chi_middleware.OapiRequestValidator(swagger))
	r.Use(httplog.RequestLogger(logger))
	r.Use(c.Recovery)
	r.Use(c.SetDBMiddleware)

	rest.HandlerFromMux(c, r)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt, os.Kill)
	defer stop()

	port := flag.Int("port", cfg.Port, "Port for HTTP server")
	flag.Parse()

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
	}

	go func() {
		<-ctx.Done()
		log.Printf("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 10)
		defer cancel()
		if err := s.Shutdown(ctx); err != nil {
			log.Printf("Error: %s\n", err)
		}
	}()

	log.Printf("Server listening on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
