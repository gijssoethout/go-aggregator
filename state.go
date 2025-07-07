package main

import (
	"github.com/gijssoethout/go-aggregator/internal/config"
	"github.com/gijssoethout/go-aggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
