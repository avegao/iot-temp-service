package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/avegao/gocondi"
	"github.com/avegao/iot-temp-service/arduino"
	"github.com/avegao/iot-temp-service/controller/default"
	pb "github.com/avegao/iot-temp-service/resource/grpc/iot_temp"
	"github.com/avegao/iot-temp-service/util"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const version = "1.0.0"

var (
	debug                       = flag.Bool("debug", false, "Print debug logs")
	iotArduinoTempServerAddress = flag.String("iot_arduino_server_addr", "iot-arduino-temp:50000", "The server address in the format of host:port")
	grcpPort                    = flag.Int("port", 50000, "gRPC Server port. Default = 50000")
	buildDate                   string
	commitHash                  string
	container                   *gocondi.Container
	server                      *grpc.Server
)

func initContainer() {
	gocondi.Initialize(initLogger())
	container = gocondi.GetContainer()

	flag.Parse()

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

	for name, value := range parameters {
		container.SetParameter(name, value)
	}
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

func initGrpc() {
	container.GetLogger().Debugf("initGrpc() - START")

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *grcpPort))

	if err != nil {
		container.GetLogger().Fatalf("failed to listen: %v", err)
	}

	container.GetLogger().Debugf("gRPC listening in %d port", *grcpPort)

	server = grpc.NewServer()
	pb.RegisterThermostatServiceServer(server, &default_controller.Controller{})
	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		container.GetLogger().Fatalf("failed to server: %v", err)
	}

	container.GetLogger().Debugf("initGrpc() - END")
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
	server.Stop()
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

	//defer powerOff()

	logger := container.GetLogger()
	logger.Infof("IoT Temperature Service started v%s (commit %s, build date %s)", container.GetStringParameter("version"), container.GetStringParameter("commit_hash"), container.GetStringParameter("build_date"))

	initGrpc()

	//repo := new(thermostat.Repository)
	//
	//thermostatObject, err := repo.FindOneById(1)
	//
	//if nil != err {
	//    logger.WithError(err).Panic()
	//}
	//
	//temp, err :=
	//thermostatObject.IsPower()
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
