package thermostat

import (
    "github.com/sirupsen/logrus"
    "fmt"
    "github.com/avegao/iot-temp-service/arduino"
    "github.com/avegao/iot-temp-service/entity/device"
    "github.com/avegao/iot-temp-service/resource/grpc"
    "errors"
)

type Thermostat struct {
    device.Device
    Auto           bool
    MinTemperature float32
}

func (thermostat Thermostat) getArduinoRequest() arduino_service.ArduinoRequest {
    if thermostat.Type != device.Arduino {
        logrus.Panicf("%s type can't covert to %s", thermostat.Type.String(), device.Arduino.String())
    }

    return arduino_service.ArduinoRequest{
        Id:   uint32(thermostat.ID),
        Name: thermostat.Name,
        Url:  fmt.Sprintf("%s:%d", thermostat.Address, thermostat.Port),
    }
}

func (thermostat Thermostat) GetTemperature() (float32, error) {
    switch thermostat.Type {
    case device.Arduino:
        return arduino.GetTemperature(thermostat.getArduinoRequest())
    default:
        return float32(0), errors.New(fmt.Sprintf("%s not support yet", thermostat.Type.String()))
    }
}

func (thermostat Thermostat) IsPower() (bool, error) {
    switch thermostat.Type {
    case device.Arduino:
        return arduino.IsPower(thermostat.getArduinoRequest())
    default:
        return false, errors.New(fmt.Sprintf("%s not support yet", thermostat.Type.String()))
    }
}

func (thermostat Thermostat) PowerOn() (bool, error) {
    switch thermostat.Type {
    case device.Arduino:
        return arduino.PowerOn(thermostat.getArduinoRequest())
    default:
        return false, errors.New(fmt.Sprintf("%s not support yet", thermostat.Type.String()))
    }
}

func (thermostat Thermostat) PowerOff() (bool, error) {
    switch thermostat.Type {
    case device.Arduino:
        return arduino.PowerOff(thermostat.getArduinoRequest())
    default:
        return false, errors.New(fmt.Sprintf("%s not support yet", thermostat.Type.String()))
    }
}
