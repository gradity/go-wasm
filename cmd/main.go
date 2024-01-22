package main

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/rs/zerolog/log"

	"gitlab.com/graditya/go-wasm/src/component"
	"gitlab.com/graditya/go-wasm/src/handler"
)

func main() {

	app.Handle("count", handler.HandleCount)

	app.Route("/", &component.Hello{})
	app.Route("/card", &component.Card{})
	app.Route("/count", &component.Counter{})
	app.Route("/countbutton", &component.CounterButtonSelf{})

	app.RunWhenOnBrowser()

	if err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "Hello Test App",
		Description: "Hello World example app",
		Resources:   app.GitHubPages("go-wasm"),
	}); err != nil {
		log.Fatal().Err(err).Send()
	}

	// http.Handle("/", &app.Handler{
	// 	Name:        "Hello",
	// 	Description: "An Hello World! example",
	// })

	// if err := http.ListenAndServe(":8000", nil); err != nil {
	// 	log.Fatal(err)
	// }

	startServer(&app.Handler{
		Name: "go-wasm",
	})
}

func startServer(app *app.Handler) {
	var port = "3300"

	log.Info().Msg("running in server mode")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app,
	}

	// Start the HTTP server in a separate goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("Failed to start server: %s", err)
		}
	}()

	log.Info().Msgf("Server listening on port %s", port)

	<-stop
	log.Info().Msg("Server shutting down...")

	// Create a context with a timeout to allow for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server and block until all connections are closed or the timeout expires
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("Failed to gracefully shut down server: %s", err)
	}

	// Perform any cleanup tasks or resource release here, if needed.

	// Close the database connection.
	// if container.DB != nil {
	// 	container.DB.Close()
	// }

	log.Info().Msg("Server gracefully stopped")
}
