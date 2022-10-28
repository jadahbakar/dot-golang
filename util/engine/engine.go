package engine

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jadahbakar/dot-golang/util/config"
	"github.com/jadahbakar/dot-golang/util/greet"
	"github.com/jadahbakar/dot-golang/util/response"
)

var logFile *os.File

func NewFiber(config *config.Config) (*fiber.App, error) {
	fa := fiber.New(fiber.Config{
		AppName:       config.App.Name,
		Prefork:       config.App.Prefork,
		CaseSensitive: config.App.CaseSensitive,
		ReadTimeout:   time.Second * time.Duration(config.App.ReadTimeOut),
		WriteTimeout:  time.Second * time.Duration(config.App.WriteTimeOut),
		ErrorHandler:  response.DefaultErrorHandler,
	})
	logFile, err := os.OpenFile(fmt.Sprintf("%s/%s", config.App.LogFolder, "fiber.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	fiberMiddleware(fa, logFile)
	greet.ConsoleGreet(config.App.Name, config.App.Version, "", config.App.Port)
	return fa, nil
}

func fiberMiddleware(fa *fiber.App, logFile *os.File) {
	ConfigLogger := logger.Config{
		Next:       nil,
		Format:     `{"pid":${pid}, "timestamp":"${time}", "status":${status}, "latency":"${latency}", "method":"${method}", "path":"${path}"}` + "\n",
		TimeFormat: "2006/Jan/02 15:04:05",
		TimeZone:   "Asia/Jakarta",
		Output:     logFile,
	}
	ConfigCORS := cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}
	ConfigCompress := compress.Config{
		Next:  nil,
		Level: compress.LevelBestSpeed,
	}
	fa.Use(
		requestid.New(),
		cors.New(ConfigCORS),
		logger.New(ConfigLogger),
		compress.New(ConfigCompress),
	)
}

// func StartFiberWithGracefulShutdown(fa *fiber.App, db *pgxpool.Pool, port int) {
func StartFiberWithGracefulShutdown(fa *fiber.App, port int) {
	// Listen from a different goroutine
	go func() {
		if err := fa.Listen(fmt.Sprintf(":%d", port)); err != nil {
			messageErr := fmt.Sprintf("Server is not running! on Port %d Reason: %v", port, err)
			log.Panicf(messageErr)
		}
	}()
	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	_ = <-c                                         // This blocks the main thread until an interrupt is received
	_ = fa.Shutdown()
	log.Printf("shutdown")
	// db.Close()
	// log.Printf("database closed")
	logFile.Close()
	log.Printf("fiberlog closed")

}
