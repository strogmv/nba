package aggregate

import (
	"context"
	"errors"
	"fmt"
	"log"
	h "net/http"
	"os"
	"os/signal"
	"time"

	na "github.com/nats-io/nats.go"
	"nba-task-main/internal/http"
	"nba-task-main/internal/http/aggregateservice"
	"nba-task-main/internal/nats"
)

// Run creates a new instance of App and runs it
func Run() error {
	app, err := NewApp()
	if err != nil {
		return fmt.Errorf("failed to init app: %w", err)
	}
	return app.Run()
}

// NewApp returns a new instance of App
func NewApp() (*App, error) {
	var app = new(App)
	var err error

	app.Config, err = NewConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to open config: %w", err)
	}

	app.NATSConn, err = nats.NewClient(app.Config.Nats)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize  nats client: %w", err)
	}

	aggregateService := aggregateservice.NewServer(app.NATSConn)
	app.HTTPServer, err = http.NewAggregateServer(app.Config.HTTP, aggregateService)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize http aggregate: %w", err)
	}

	return app, nil
}

// App contains all what needs to run the aggregate
type App struct {
	Config     *Config
	HTTPServer *h.Server
	NATSConn   *na.Conn
}

// Run runs App and waits ending of app life-cycle
func (app *App) Run() error {
	var errc = app.Start()

	var quit = make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	select {
	case <-quit:
		log.Println("caught os signal")
	case err := <-errc:
		log.Printf("caught error: %s", err)
	}

	log.Println("trying to shutdown aggregate")

	return app.Shutdown(context.TODO())
}

// Start runs App and doesn't wait
func (app *App) Start() <-chan error {
	var errc = make(chan error, 1)

	go func() {
		log.Println("http aggregate has been started")
		if err := app.HTTPServer.ListenAndServe(); err != nil && !errors.Is(err, h.ErrServerClosed) {
			errc <- err
		}
	}()

	return errc
}

// Shutdown can be run to clean up all what was run
func (app *App) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := app.HTTPServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown http aggregate")
	}

	app.NATSConn.Close()

	log.Println("aggregate has been shutdown")
	return nil
}
