package utils

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func StartServer(app *fiber.App) {

	// Listen o different goroutine
	go func(){
		if err := app.Listen(os.Getenv("SERVER_URL")); err != nil {
			log.Panicf("Ooops... Server is not running! Reason: %v", err)
		} 
	}()

	closeConn := make(chan os.Signal, 1)
	signal.Notify(closeConn, os.Interrupt, syscall.SIGTERM)

	<- closeConn // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down")
	_= app.Shutdown()

	fmt.Println("Running clean up tasks")
	// clean up tasks go here
	// db.Close()
	// redisConn.Close()

	fmt.Println("Server was shutdown successfully")


}