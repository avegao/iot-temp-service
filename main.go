package main

import (
    "flag"
    "github.com/sirupsen/logrus"
    "os"
    "os/signal"
    "syscall"
    deviceEntity "github.com/avegao/iot-temp-service/entity/device"
    "github.com/avegao/iot-temp-service/arduino"
)

const (
    version = "0.1.0"
)

var (
    debug                       = flag.Bool("debug", false, "Print debug logs")
    iotArduinoTempServerAddress = flag.String("iot_arduino_server_addr", "iot-arduino-temp:50000", "The server address in the format of host:port")
    //port                        = flag.Int("port", 50000, "gRPC server port")
    //tls                         = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
    //caFile                      = flag.String("ca", "", "The file containning the CA root cert file")
    buildDate   string
    commmitHash string
)

func initLogger() {
    logLevel := logrus.InfoLevel

    if *debug {
        logLevel = logrus.DebugLevel
    }

    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetFormatter(&logrus.TextFormatter{})
    logrus.SetLevel(logLevel)
}

func handleInterrupt() {
    gracefulStop := make(chan os.Signal)
    signal.Notify(gracefulStop, syscall.SIGTERM)
    signal.Notify(gracefulStop, syscall.SIGINT)
    go func() {
        <-gracefulStop
        powerOff()
    }()
}

func powerOff() {
    logrus.Infof("Shutting down...")
    arduino.CloseConnection()
    os.Exit(0)
}

func main() {
    flag.Parse()
    handleInterrupt()
    initLogger()

    arduino.Address = *iotArduinoTempServerAddress

    logrus.Infof("IoT Temperature Service started v%s (commit %s, build date %s)", version, commmitHash, buildDate)

    arduinoDevice, err := deviceEntity.New().SetTypeArduino().SetId(1).SetName("Arduino").SetAddress("192.168.1.163").SetPort(80).Build()

    if nil != err {
        logrus.WithError(err).Panic()
    }

    temp, err := arduinoDevice.GetTemperature()

    if nil != err {
        logrus.WithError(err).Panicf("Error getting temp")
    }

    logrus.Infof("Temperature = %f", temp)

    power, err := arduinoDevice.IsPower()

    if nil != err {
        logrus.WithError(err).Panicf("Error getting temp")
    }

    logrus.Infof("Power = %b", power)

    //_, err = arduinoDevice.PowerOn()
    //
    //if nil != err {
    //    logrus.WithError(err).Panicf("Error getting temp")
    //}
    //
    //time.Sleep(time.Second * 15)
    //
    //_, err = arduinoDevice.PowerOff()
    //
    //if nil != err {
    //    logrus.WithError(err).Panicf("Error getting temp")
    //}

    powerOff()
}
