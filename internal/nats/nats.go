package nats

import (
	"fmt"
	na "github.com/nats-io/nats.go"
	"log"
)

type Request struct {
	Body interface{}
}
type Responce struct {
	Code  int
	Data  interface{}
	Error string
}

type Config struct {
	Host string `yaml:"host"`
}

func NewClient(cfg Config) (*na.Conn, error) {
	nc, err := na.Connect(cfg.Host)
	if err != nil {
		log.Println("NATS Connect", err)
		return nil, fmt.Errorf("failed to initialize NATS client: %w", err)
	}

	return nc, nil
}
