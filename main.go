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

const version = "1.0.0"

var (
    debug                       = flag.Bool("debug", false, "Print debug logs")
    iotArduinoTempServerAddress = flag.String("iot_arduino_server_addr", "iot-arduino-temp:50000", "The server address in the format of host:port")
    buildDate                   string
    commitHash                  string
    container                   *util.Container
)

func initContainer() {
    container = util.GetContainer()

    flag.Parse()

    logger := initLogger()
    container.SetLogger(logger)

    db := initPgsql()
    container.SetDefaultDatabase(db)
    container.SetDatabase("pgsql", db)

    parameters := map[string]interface{}{
        "build_date":                      buildDate,
        "debug":                           *debug,
        "commit_hash":                     commitHash,
        "iot_arduino_temp_server_address": *iotArduinoTempServerAddress,
        "version":                         version,
    }

    container.SetParameters(parameters)
}

func initLogger() *logrus.Logger {
    logLevel := logrus.InfoLevel

    if *debug {
        logLevel = logrus.DebugLevel
    }

    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetFormatter(&logrus.TextFormatter{})

    log := logrus.New()
    log.SetLevel(logLevel)

    return log
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

func closeDatabases() {
    for _, database := range container.GetDatabases() {
        database.Close()
    }
}

func powerOff() {
    container.GetLogger().Infof("Shutting down...")
    arduino.CloseConnection()
    closeDatabases()
    os.Exit(0)
}

func initPgsql() *sql.DB {
    const logTag = "initPgsql() ->"
    logger := container.GetLogger()
    logger.Debugf("%s START", logTag)

    host := util.GetStringParameter("PGSQL_HOST", "localhost")
    user := util.GetStringParameter("PGSQL_USERNAME", "iot")
    password := util.GetStringParameter("PGSQL_PASSWORD", "iot")
    dbName := util.GetStringParameter("PGSQL_DATABASE", "iot")
    connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)

    db, err := sql.Open("postgres", connection)

    if nil != err {
        logger.Panic(err)
    }

    logger.Debugf("%s Database connected", logTag)
    logger.Debugf("%s END", logTag)

    return db
}

func main() {
    initContainer()
    handleInterrupt()

    defer powerOff()

    logger := container.GetLogger()
    logger.Infof("IoT Temperature Service started v%s (commit %s, build date %s)", container.GetParameter("version"), container.GetParameter("commit_hash"), container.GetParameter("build_date"))

    repo := new(thermostat.Repository)

    thermostatObject, err := repo.FindOneById(1)

    if nil != err {
        logger.WithError(err).Panic()
    }

    //temp, err :=
    thermostatObject.PowerOn()
    //
    //if nil != err {
    //    logger.WithError(err).Panicf("Error getting temp")
    //}
    //
    //logger.Infof("Temp = %f", temp)

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
