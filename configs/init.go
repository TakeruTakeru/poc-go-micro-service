package configs

import (
	"fmt"
	"github.com/TakeruTakeru/poc-go-micro-service/pkg/logger"
	"github.com/joho/godotenv"
	"os"
)

const (
	PROJECT_NAME = "poc-go-micro-service"
)

var ROOT_PATH = fmt.Sprintf("%s/src/github.com/TakeruTakeru/%s", os.Getenv("GOPATH"), PROJECT_NAME)

func InitCommonEnv() (err error) {
	err = godotenv.Load(ROOT_PATH + "/.env")
	if err != nil {
		logger.Errorf("Failed to load .env file")
	}
	logger.Printf("Initialize common environment.")
	return
}

func InitDevEnv() (err error) {
	err = godotenv.Load(ROOT_PATH + "/configs/devconf/.env")
	if err != nil {
		logger.Errorf("Failed to load .env file")
	}
	logger.Printf("Initialize dev environment.")
	return
}

func InitTestEnv() (err error) {
	err = godotenv.Load(ROOT_PATH + "/configs/testconf/.env")
	if err != nil {
		logger.Errorf("Failed to load .env file")
	}
	logger.Printf("Initialize test environment")
	return
}

func init() {
	err := InitCommonEnv()
	if err != nil {
		logger.Errorf(err.Error())
	}
	var env string
	env = os.Getenv("GOENV")
	switch env {
	case "test":
		err = InitTestEnv()
	case "prod":
		logger.Panicf("no impl")
	default:
		err = InitDevEnv()
	}
	if err != nil {
		logger.Errorf(err.Error())
	}
}
