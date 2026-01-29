package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/Joshua-Pok/BlogAggregator/internal/database"
	"github.com/google/uuid"
	"log"
	"time"
)

type Command struct {
	Name      string
	Arguments []string
}

type State struct {
	Cfg *Config
	Db  *database.Queries
}

type Commands struct {
	Commands_to_handler map[string]func(*State, Command) error // mapping of commands to their appropriate handler function
}

// runs a given command with the provided state if it exists
func (c *Commands) Run(s *State, cmd Command) error {
	// value, ok := myMap[key]
	handler, exists := c.Commands_to_handler[cmd.Name]
	if !exists {

		log.Fatalf("No handler for this command")
	}

	return handler(s, cmd)
}

// registers a new handler function for a command name
func (c *Commands) Register(name string, f func(*State, Command) error) {

	_, exists := c.Commands_to_handler[name]
	if exists {
		fmt.Printf("Handler for %v already exists!", name)
	}

	c.Commands_to_handler[name] = f

}

func HandlerLogin(s *State, cmd Command) error {

	if len(cmd.Arguments) == 0 {
		return errors.New("No arguments passed")

	}

	username := cmd.Arguments[0]
	SetUser(s.Cfg, username)

	fmt.Printf("User has been set to %v", username)
	return nil

}

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return errors.New("No arguments passed")
	}

	name := cmd.Arguments[0]
	_, err := s.Db.GetUser(context.Background(), name)
	if err == nil {
		fmt.Printf("User %v already exists!", name)
		return nil
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	user, err := s.Db.CreateUser(context.Background(), params)
	if err != nil {
		log.Fatalf("Error registering user: %v", err)
	}

	SetUser(s.Cfg, name)

	fmt.Printf("User %v was created, %v+b", name, user)
	return nil

}

func Reset(s *State, cmd Command) error {
	err := s.Db.TruncateUsers(context.Background())
	if err != nil {
		log.Fatalf("Error resetting users: %v", err)
	}
	return nil
}

func users(s *State, cmd Command) ([]database.User, error) {
	users, err := s.Db.GetAllUsers(context.Background())
	if err != nil {
		log.Fatalf("Error retrieving all users: %v", err)
	}

	for _, user := range users {
		if user.Name == s.Cfg.Current_user_name {
			fmt.Printf("%v, (current) \n", user.Name)
		} else {
			fmt.Println(user.Name)
		}
	}

	return users, nil
}

func addfeed(name string, url string) error {

}
