package util

import (
    "flag"
    "github.com/sirupsen/logrus"
    "database/sql"
)

const Version = "0.1.0"

var (
    Debug                       = flag.Bool("debug", false, "Print debug logs")
    IotArduinoTempServerAddress = flag.String("iot_arduino_server_addr", "iot-arduino-temp:50000", "The server address in the format of host:port")
    Log                         *logrus.Logger
    Pgsql                       *sql.DB
)
