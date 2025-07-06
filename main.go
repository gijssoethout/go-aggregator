package main

import (
	"fmt"

	"github.com/gijssoethout/go-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	cfg.SetUser("Gijs")
	newCfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newCfg.DBURL)
	fmt.Println(newCfg.CurrentUserName)
}
