package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

// go does not read ~ inthe file path we need to manually resolve the home directory
func getConfigFilePath() (string, error) {

	home_dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Home directory not found: %v", err)
	}

	return home_dir + "/" + configFileName, nil

}

func Read() Config { //capital letter to export

	configFile, err := getConfigFilePath()
	if err != nil {
		log.Fatalf("Unable to retrieve config file path: %v", err)
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Failed to find config file: %v", err)
	}

	//declare a config instance
	var config Config

	if err = json.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}
	// impt: unmarshal needs a pointer

	return config

}

func write(cfg []byte) error {

	fileName, err := getConfigFilePath()
	if err != nil {
		log.Fatalf("write: unable to get filepath: %v", err)
	}

	err = os.WriteFile(fileName, cfg, 0644)

	return nil
}

// in go function arguments are copied unless we pass a pointer
// important to note marshal only converts a go value into []bytes (JSON in memory)
// hence we need a seperate write function to wrtie it back to the file

func SetUser(config *Config, username string) {

	config.Current_user_name = username

	data, err := json.Marshal(config)
	if err != nil {
		log.Fatalf("Failed to marshal config back: %v", err)
	}

	write(data)

}
