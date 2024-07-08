package main

import (
	"log"

	stats "nba-task-main/internal/app/stats"
)

func main() {
	if err := stats.Run(); err != nil {
		log.Fatal(err)
	}
}
