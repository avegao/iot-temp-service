package thermostat

import (
    "github.com/sirupsen/logrus"
    "github.com/avegao/iot-temp-service/entity/device"
    "github.com/avegao/iot-temp-service/util"
    "errors"
    "fmt"
)

type ThermostatBuilderInterface interface {
    Build() Thermostat
    SetId(id uint64) thermostatBuilder
    SetName(name string) thermostatBuilder
    SetAddress(address string) thermostatBuilder
    SetPort(port int) thermostatBuilder
    SetAuto(auto bool) thermostatBuilder
    SetMinTemperature(temperature float32) thermostatBuilder
}

type thermostatBuilder struct {
    ThermostatBuilderInterface
    thermostat Thermostat
}

func (builder thermostatBuilder) SetId(id uint64) thermostatBuilder {
    builder.thermostat.ID = id

    return builder
}

func (builder thermostatBuilder) SetName(name string) thermostatBuilder {
    builder.thermostat.Name = name

    return builder
}

func (builder thermostatBuilder) SetAddress(address string) thermostatBuilder {
    builder.thermostat.Address = address

    return builder
}

func (builder thermostatBuilder) SetPort(port int) thermostatBuilder {
    builder.thermostat.Port = port

    return builder
}

func (builder thermostatBuilder) SetTypeArduino() thermostatBuilder {
    builder.thermostat.Type = device.Arduino

    return builder
}

func (builder thermostatBuilder) SetTypeRaspberryPi() thermostatBuilder {
    builder.thermostat.Type = device.RaspberryPi

    return builder
}

func (builder thermostatBuilder) SetAuto(auto bool) thermostatBuilder {
    builder.thermostat.Auto = auto

    return builder
}

func (builder thermostatBuilder) SetMinTemperature(temperature float32) thermostatBuilder {
    builder.thermostat.MinTemperature = temperature

    return builder
}

func (builder thermostatBuilder) Build() (Thermostat, error) {
    if builder.thermostat.ID < 0 {
        logrus.Warnf("ID must be >= 0. Setting to 0. This result on a new device.")

        builder.thermostat.ID = 0
    }

    if util.Empty(builder.thermostat.Name) {
        return *new(Thermostat), errors.New("device name is required")
    }

    if util.Empty(builder.thermostat.Address) {
        return *new(Thermostat), errors.New("device address is required")
    }

    if 0 == builder.thermostat.Port {
        return *new(Thermostat), errors.New("device port is required")
    }

    if 0 == builder.thermostat.Type {
        return *new(Thermostat), errors.New("device type is required")
    } else if builder.thermostat.Type != device.Arduino {
        logrus.Panicf("%s not support yet", builder.thermostat.Type)
    }

    if float32(0) >= builder.thermostat.MinTemperature {
        return *new(Thermostat), errors.New(fmt.Sprintf("min temperature must be > 0. You set %f", builder.thermostat.MinTemperature))
    }

    return builder.thermostat, nil
}

func New() thermostatBuilder {
    return thermostatBuilder{}
}
