package service

import (
	"afet-yardim-twitter-bot/config"
	"afet-yardim-twitter-bot/pkg/handler"
	"afet-yardim-twitter-bot/pkg/router"
	"afet-yardim-twitter-bot/pkg/service"
	"bytes"
	"context"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"
)

var (
	cfg    config.Config
	logger logrus.Logger
)

func Run() {

	initConfig()
	initLogger()

	twitterClient, err := initTwitterClient()
	if err != nil {
		logger.Errorf(err.Error())
		return
	}

	svc := service.NewBotService(&logger, twitterClient)
	handlers := handler.New(svc)
	server := initHTTPHandler(handlers)

	// start the server
	go func() {
		logger.Println("Starting server on", cfg.Server.HttpAddress)

		err := server.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	logger.Println("Got signal: ", sig)

	// gracefully shutdown the server
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = server.Shutdown(ctx)
}

func initHTTPHandler(handlers handler.Handlers) *http.Server {

	server := &http.Server{
		Addr:         cfg.Server.HttpAddress, // configure the bind address
		ReadTimeout:  50 * time.Second,       // max time to read request from the client
		WriteTimeout: 100 * time.Second,      // max time to write response to the client
		IdleTimeout:  12 * time.Second,       // max time for connections using TCP Keep-Alive
		Handler:      router.NewRouter(handlers),
	}

	return server
}

func initConfig() {
	err := envconfig.Process("", &cfg)
	if err == nil {
		logger.Info("configs are loaded from environment. no need to load .env file")
		return
	}

	envFile := func() string {
		ef := os.Getenv("ENV_FILE")
		if ef == "" {
			return ".env"
		}
		return ef
	}()

	err = godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("unable to load env file: %v error: %e", envFile, err)
	}

	err = envconfig.Process("", &cfg)

	if err != nil {
		logger.Fatalf(err.Error())
	}
}

func initLogger() {
	logger = *logrus.New()

	logger.Level = logrus.InfoLevel
	logger.Formatter = &formatter{}

	logger.SetReportCaller(true)
}

type formatter struct {
	prefix string
}

func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var sb bytes.Buffer

	var newLine = "\n"
	if runtime.GOOS == "windows" {
		newLine = "\r\n"
	}

	sb.WriteString(strings.ToUpper(entry.Level.String()))
	sb.WriteString(" ")
	sb.WriteString(entry.Time.Format(time.RFC3339))
	sb.WriteString(" ")
	sb.WriteString(f.prefix)
	sb.WriteString(entry.Message)
	sb.WriteString(newLine)

	return sb.Bytes(), nil
}
