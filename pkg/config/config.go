package config

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ConfigSpec struct {
	Port              string
	AuthorizedKeys    []string
	DiscordWebhookURL string
}

var Config = ConfigSpec{
	Port:              "8080",
	AuthorizedKeys:    []string{},
	DiscordWebhookURL: "",
}

func LoadConfig() {
	// try to laod from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Unable to load .env file continuing without it")
	}

	Config.Port = os.Getenv("PORT")
	if Config.Port == "" {
		Config.Port = "8080"
	}

	Config.DiscordWebhookURL = os.Getenv("DISCORD_WEBHOOK_URL")

	// Read authorized keys from configuration file
	keysFile := os.Getenv("KEYS_FILE")
	if keysFile == "" {
		keysFile = "keys.txt"
	}

	// Check if the file exists
	if _, err := os.Stat(keysFile); os.IsNotExist(err) {
		// File does not exist, log an error and exit
		fmt.Println("Keys file does not exist:", keysFile)
		os.Exit(1)
	}

	// Read the file
	file, err := os.Open(keysFile)
	if err != nil {
		fmt.Println("Error opening keys file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the file contents
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line[0] == '#' {
			continue
		}
		Config.AuthorizedKeys = append(Config.AuthorizedKeys, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading keys file:", err)
		os.Exit(1)
	}
}
