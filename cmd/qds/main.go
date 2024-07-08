package main

import (
	"log"

	"nba-task-main/internal/app/qds"
)

func main() {
	if err := qds.Run(); err != nil {
		log.Fatal(err)
	}
}
