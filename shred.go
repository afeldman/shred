package main

import (
	"shred/cmd"

	log "github.com/sirupsen/logrus"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Unexpected error: %v", r)
		}
	}()
	cmd.Execute()
}
