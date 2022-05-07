package configs

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	readTimeoutSecondsCount, err := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	if err != nil {
		log.Fatal("Read timeout could not be set.")
	}

	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount), 
	}
}