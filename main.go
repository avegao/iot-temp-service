package main

import (
    "flag"
    "github.com/sirupsen/logrus"
)

var (
    debug = false
)

func initLogger() {
    logLevel := logrus.InfoLevel

    if debug {
        logLevel = logrus.DebugLevel
    }

    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetFormatter(&logrus.TextFormatter{})
    logrus.SetLevel(logLevel)
}

func getParameters() {
    debug = *flag.Bool("debug", false, "Print debug logs")

    flag.Parse()
}

func main() {
    getParameters()
    initLogger()

    logrus.Infof("IoT Arduino Temperature Service started")
}
