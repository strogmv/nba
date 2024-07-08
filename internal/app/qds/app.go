package qds

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"nba-task-main/internal/entity"
	"nba-task-main/internal/postgres"
	"os"
	"os/signal"
	"time"

	na "github.com/nats-io/nats.go"
	"nba-task-main/internal/nats"
)

// Run creates a new instance of App and runs it
func Run() error {
	log.Println("run qds app")
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
	log.Println("app.Config.DSN", app.Config.DSN)
	db, err := postgres.NewDB(app.Config.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize  db: %w", err)
	}

	app.Repository = postgres.NewRepository(db)
	return app, nil
}

// App contains all what needs to run the api
type App struct {
	Config     *Config
	Repository *postgres.Repository
	NATSConn   *na.Conn
}

// Run runs App and waits ending of app life-cycle
func (app *App) Run() error {
	var quit = make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	_, err := app.NATSConn.Subscribe(entity.AddPlayerNATSHandler, app.AddPlayerHandler)
	if err != nil {
		log.Printf("nats start error %v", err)
		return app.Shutdown(context.TODO())
	}
	_, err = app.NATSConn.Subscribe(entity.AddTeamNATSHandler, app.AddTeamHandler)
	if err != nil {
		log.Printf("nats start error %v", err)
		return app.Shutdown(context.TODO())
	}
	_, err = app.NATSConn.Subscribe(entity.AddStatNATSHandler, app.AddStatHandler)
	if err != nil {
		log.Printf("nats start error %v", err)
		return app.Shutdown(context.TODO())
	}
	_, err = app.NATSConn.Subscribe(entity.GetPlayerNATSHandler, app.GetPlayerHandler)
	if err != nil {
		log.Printf("nats start error %v", err)
		return app.Shutdown(context.TODO())
	}
	_, err = app.NATSConn.Subscribe(entity.GetTeamNATSHandler, app.GetTeamHandler)
	if err != nil {
		log.Printf("nats start error %v", err)
		return app.Shutdown(context.TODO())
	}
	<-quit
	log.Println("caught os signal")

	log.Println("trying to shutdown api")

	return app.Shutdown(context.TODO())
}

// Shutdown can be run to clean up all what was run
func (app *App) Shutdown(ctx context.Context) error {
	_, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	app.NATSConn.Close()

	log.Println("api has been shutdown")
	return nil
}

func (app *App) AddPlayerHandler(msg *na.Msg) {
	var re entity.Player
	err := json.Unmarshal(msg.Data, &re)
	if err != nil {
		log.Printf("nats start error %v", err)
	}

	err = app.Repository.AddPlayer(re)

	if err != nil {
		errorResp := nats.Responce{Code: 400, Error: err.Error()}
		b, err := json.Marshal(errorResp)
		if err != nil {
			log.Printf("nats resp error %v", err)
		}
		err = msg.Respond(b)
		if err != nil {
			log.Printf("nats resp error %v", err)
		}
	}
	okResp := nats.Responce{Code: 200}
	b, err := json.Marshal(okResp)
	if err != nil {
		log.Printf("nats resp error %v", err)
	}
	err = msg.Respond(b)
	if err != nil {
		log.Printf("nats resp error %v", err)
	}
}

func (app *App) AddTeamHandler(msg *na.Msg) {
	var re entity.Team
	err := json.Unmarshal(msg.Data, &re)
	if err != nil {
		log.Printf("nats start error %v", err)
	}
	err = app.Repository.AddTeam(re)
	if err != nil {
		log.Println("Repository.AddTeam", err)
		errorResp := nats.Responce{Code: 400, Error: err.Error()}
		b, err := json.Marshal(errorResp)
		if err != nil {
			log.Printf("nats resp error %v", err)
		}
		err = msg.Respond(b)
		if err != nil {
			log.Printf("nats resp error %v", err)
		}
	}
	okResp := nats.Responce{Code: 200}
	b, err := json.Marshal(okResp)
	if err != nil {
		log.Println("Repository.AddTeam3", err)
		log.Printf("nats resp error %v", err)
	}
	err = msg.Respond(b)
	if err != nil {
		log.Printf("nats resp error %v", err)
	}
}
func (app *App) AddStatHandler(msg *na.Msg) {
	var re entity.Stat
	err := json.Unmarshal(msg.Data, &re)
	if err != nil {
		log.Printf("nats start error %v", err)
	}
	err = app.Repository.AddStat(re)

	if err != nil {
		errorResp := nats.Responce{Code: 400, Error: err.Error()}
		b, err := json.Marshal(errorResp)
		if err != nil {
			log.Printf("nats resp error %v", err)
		}
		err = msg.Respond(b)
		if err != nil {
			log.Printf("nats resp error %v", err)
		}
	}
	okResp := nats.Responce{Code: 200}
	b, err := json.Marshal(okResp)
	if err != nil {
		log.Printf("nats resp error %v", err)
	}
	err = msg.Respond(b)
	if err != nil {
		log.Printf("nats resp error %v", err)
	}
}
func (app *App) GetPlayerHandler(msg *na.Msg) {
	var playerID int
	err := json.Unmarshal(msg.Data, &playerID)
	if err != nil {
		log.Printf("nats start error %v", err)
	}
	resp, err := app.Repository.PlayerAverage(playerID)

	if err != nil {
		errorResp := nats.Responce{Code: 400, Error: err.Error()}
		b, err := json.Marshal(errorResp)
		if err != nil {
			log.Printf("nats resp error %v", err)
		}
		err = msg.Respond(b)
		if err != nil {
			log.Printf("nats resp error %v", err)
		}
	}
	b, err := json.Marshal(resp)
	if err != nil {
		log.Printf("nats resp error %v", err)
	}
	err = msg.Respond(b)
	if err != nil {
		log.Printf("nats resp error %v", err)
	}
}
func (app *App) GetTeamHandler(msg *na.Msg) {
	var teamID int
	err := json.Unmarshal(msg.Data, &teamID)
	if err != nil {
		log.Printf("nats start error %v", err)
	}
	resp, err := app.Repository.TeamAverage(teamID)

	if err != nil {
		errorResp := nats.Responce{Code: 400, Error: err.Error()}
		b, err := json.Marshal(errorResp)
		if err != nil {
			log.Printf("nats resp error %v", err)
		}
		err = msg.Respond(b)
		if err != nil {
			log.Printf("nats resp error %v", err)
		}
	}
	b, err := json.Marshal(resp)
	if err != nil {
		log.Printf("nats resp error %v", err)
	}
	err = msg.Respond(b)
	if err != nil {
		log.Printf("nats resp error %v", err)
	}
}
