package device

import (
    "github.com/sirupsen/logrus"
    "github.com/avegao/iot-temp-service/util"
    "github.com/avegao/iot-temp-service/resource/grpc"
    "fmt"
    "github.com/avegao/iot-temp-service/arduino"
    "errors"
)

type Type int

func (deviceType Type) String() string {
    switch deviceType {
    case Arduino:
        return "ARDUINO"
    case RaspberryPi:
        return "RASPBERRY_PI"
    default:
        return "ANOTHER"
    }
}

const (
    Arduino     Type = 1
    RaspberryPi Type = 2
)

type Device struct {
    util.Entity
    Name    string
    Address string
    Port    int
    Type    Type
}

func (device Device) getArduinoRequest() arduino_service.ArduinoRequest {
    if device.Type != Arduino {
        logrus.Panicf("%s type can't covert to %s", device.Type.String(), Arduino.String())
    }

    return arduino_service.ArduinoRequest{
        Id:   uint32(device.ID),
        Name: device.Name,
        Url:  fmt.Sprintf("%s:%d", device.Address, device.Port),
    }
}

func (device Device) GetTemperature() (float32, error) {
    switch device.Type {
    case Arduino:
        return arduino.GetTemperature(device.getArduinoRequest())
    default:
        return float32(0), errors.New(fmt.Sprintf("%s not support yet", device.Type.String()))
    }
}

func (device Device) IsPower() (bool, error) {
    switch device.Type {
    case Arduino:
        return arduino.IsPower(device.getArduinoRequest())
    default:
        return false, errors.New(fmt.Sprintf("%s not support yet", device.Type.String()))
    }
}

func (device Device) PowerOn() (bool, error) {
    switch device.Type {
    case Arduino:
        return arduino.PowerOn(device.getArduinoRequest())
    default:
        return false, errors.New(fmt.Sprintf("%s not support yet", device.Type.String()))
    }
}

func (device Device) PowerOff() (bool, error) {
    switch device.Type {
    case Arduino:
        return arduino.PowerOff(device.getArduinoRequest())
    default:
        return false, errors.New(fmt.Sprintf("%s not support yet", device.Type.String()))
    }
}
