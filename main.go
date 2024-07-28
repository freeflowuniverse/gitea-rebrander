package main

import (
	"log"

	"github.com/freeflowuniverse/gitea-rebrander/cmd"
	config "github.com/freeflowuniverse/gitea-rebrander/internal"
)

func main() {

	config, err := config.ReadConfig()

	if err != nil {
		log.Fatalf("failed to read config %v", err)
	}

	transporter := cmd.NewTransporter(config)

	if err := transporter.Transport(); err != nil {
		log.Fatalf("failed to transport %v", err)
	}
}
