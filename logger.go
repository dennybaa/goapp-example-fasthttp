package main

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// var logfile *zap.SugaredLogger
var logger *zap.SugaredLogger

func initLogger() {
	var config zap.Config
	level, err := zap.ParseAtomicLevel(appConfig.LogLevel)

	if err != nil {
		log.Panic(err)
	}

	// use efficient encoder for console only, since
	// a hard-fixed format is required for the file
	config = zap.NewProductionConfig()
	config.Development = false // we don't need extended stackrace with every error message
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.Level = level

	// set the default logger
	l := zap.Must(config.Build())
	logger = l.Sugar() // switch to sugar mode
}
