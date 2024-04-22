package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func Init() {
	initEnv()

}

func initEnv() {

	var appMode string

	flag.StringVar(&appMode, "mode", "dev", "Set the application mode (dev, prod)")
	flag.Parse()

	if appMode != "dev" && appMode != "prod" {
		log.Fatalf("Invalid mode: %s. Must be 'dev' or 'prod'", appMode)
	}

	log.Println("APP Mode: ", appMode)

	path, _ := filepath.Abs(fmt.Sprintf("./.env.%s", appMode))
	log.Println("path env:", path)
	err := godotenv.Load(path)

	if err != nil {
		log.Fatalf("Error loading environment: %v", err)
	}
	log.Println("env APPLICATION_NAME: ", os.Getenv("APPLICATION_NAME"))
}
