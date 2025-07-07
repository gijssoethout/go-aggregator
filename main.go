package main

import (
	"log"
	"os"

	"github.com/gijssoethout/go-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{
		cfg: &cfg,
	}

	c := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	c.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}
	if err := c.run(programState, cmd); err != nil {
		log.Fatal(err)
	}
}
