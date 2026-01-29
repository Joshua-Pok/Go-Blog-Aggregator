package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Joshua-Pok/BlogAggregator/internal/config"
	"github.com/Joshua-Pok/BlogAggregator/internal/database"
	"github.com/Joshua-Pok/BlogAggregator/internal/rss"
	_ "github.com/lib/pq" // _ tells go that we import it for side effects so we dont get compile errors
)

func main() {

	configStruct := config.Read()

	// config.SetUser(&configStruct, "Joshua")
	var configState config.State

	configState.Cfg = &configStruct

	// Load DB

	dbURL := configStruct.Db_url

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to Db: %v", err)
	}

	dbQueries := database.New(db)

	configState.Db = dbQueries

	var commands2Handlers config.Commands
	commands2Handlers.Commands_to_handler = make(map[string]func(*config.State, config.Command) error)
	commands2Handlers.Register("login", config.HandlerLogin)
	commands2Handlers.Register("register", config.HandlerRegister)
	commands2Handlers.Register("addfeed", config.HandlerAddFeed)

	allArgs := os.Args
	userArgs := allArgs[1:] //skip the command arg

	if len(userArgs) < 2 {
		log.Fatalf("Need minimum 2 arguments")
		return
	}

	cmd := config.Command{
		Name:      userArgs[0],
		Arguments: userArgs[1:],
	}

	err = commands2Handlers.Run(&configState, cmd)

	if err != nil {
		log.Fatal(err)
	}

	// configStruct = config.Read()
	//
	// fmt.Printf("%+v", configStruct)

}
