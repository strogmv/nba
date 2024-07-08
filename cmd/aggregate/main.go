package main

import (
	"log"

	"nba-task-main/internal/app/aggregate"
)

func main() {
	if err := aggregate.Run(); err != nil {
		log.Fatal(err)
	}
}
