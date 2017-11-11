package main

import (
    "database/sql"
    "flag"
    "fmt"
    "github.com/avegao/iot-temp-service/arduino"
    "github.com/avegao/iot-temp-service/util"
    "github.com/sirupsen/logrus"
    "os"
    "os/signal"
    "syscall"
    _ "github.com/lib/pq"
    "github.com/avegao/iot-temp-service/entity/device/thermostat"
)

var (
    buildDate  string
    commitHash string
)

func initLogger() {
    logLevel := logrus.InfoLevel

    if *util.Debug {
        logLevel = logrus.DebugLevel
    }

    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetFormatter(&logrus.TextFormatter{})

    util.Log = logrus.New()
    util.Log.SetLevel(logLevel)
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
    util.Log.Infof("Shutting down...")
    arduino.CloseConnection()
    util.Pgsql.Close()
    os.Exit(0)
}

func initPgsql() {
    const logTag = "initPgsql() ->"
    util.Log.Debugf("%s START", logTag)

    host := util.GetStringParameter("PGSQL_HOST", "localhost")
    user := util.GetStringParameter("PGSQL_USERNAME", "iot")
    password := util.GetStringParameter("PGSQL_PASSWORD", "iot")
    dbName := util.GetStringParameter("PGSQL_DATABASE", "iot")
    connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)

    db, err := sql.Open("postgres", connection)

    if nil != err {
        util.Log.Panic(err)
    }

    util.Log.Debugf("%s Database connected", logTag)

    util.Pgsql = db

    util.Log.Debugf("%s END", logTag)
}

func main() {
    flag.Parse()
    handleInterrupt()
    initLogger()
    initPgsql()
    defer powerOff()

    util.Log.Infof("IoT Temperature Service started v%s (commit %s, build date %s)", util.Version, commitHash, buildDate)

    thermostatObject, err := new(thermostat.Repository).FindOneById(1)

    if nil != err {
        util.Log.WithError(err).Panic()
    }

    power, err := thermostatObject.IsPower()

    if nil != err {
        logrus.WithError(err).Panicf("Error getting temp")
    }

    logrus.Infof("Power = %s", power)

    //arduinoDevice, err := thermostat.New().SetTypeArduino().SetId(1).SetName("Arduino").SetAddress("192.168.1.163").SetPort(80).SetAuto(true).SetMinTemperature(float32(17)).Build()
    //
    //if nil != err {
    //    util.Log.WithError(err).Panic()
    //}
    //
    //temp, err := arduinoDevice.GetTemperature()
    //
    //if nil != err {
    //    util.Log.WithError(err).Panicf("Error getting temp")
    //}
    //
    //util.Log.Infof("Temperature = %f", temp)

    //power, err := arduinoDevice.IsPower()
    //
    //if nil != err {
    //    logrus.WithError(err).Panicf("Error getting temp")
    //}
    //
    //logrus.Infof("Power = %b", power)
    //
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
}
