package internal

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

/*
LoadEnvVariables loads environment variables from the .env file using godotenv.
It checks if the .env file exists and loads its content. If successful, it continues execution.
If the file cannot be loaded, it logs a fatal error and exits the application.
*/
func LoadEnvVariables() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if err := checkEnvVariablesArePopulated(); err != nil {
		return err
	}

	return nil
}

/*
checkEnvVariablesArePopulated checks if each required environment variable is set before continuing the application execution.
Iterates over a list of required environment variables and checks if each one is set.
If any required variable is not set, it logs a fatal error and exits the application.
*/
func checkEnvVariablesArePopulated() error {
	requiredEnvVars := []string{"ENVIRONMENT", "DB_URL", "DB_NAME", "USERS_COLLECTION", "USER_SESSION_COLLECTION", "HABITS_COLLECTION", "ENV", "PORT", "SITE_URL", "API_NAME", "APP_VERSION", "API_VERSION", "LOG_VERBOSITY", "JWT_SECRET", "DEBUG"}

	for _, envVal := range requiredEnvVars {
		if value := os.Getenv(envVal); value == "" {
			return fmt.Errorf("environment variable '%s' is not set", envVal)
		}
	}

	return nil
}
