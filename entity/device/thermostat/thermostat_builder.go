package thermostat

import (
    "github.com/sirupsen/logrus"
    "github.com/avegao/iot-temp-service/entity/device"
    "github.com/avegao/iot-temp-service/util"
    "errors"
    "fmt"
)

type BuilderInterface interface {
    Build() Thermostat
    SetId(id uint64) Builder
    SetName(name string) Builder
    SetAddress(address string) Builder
    SetPort(port int) Builder
    SetAuto(auto bool) Builder
    SetMinTemperature(temperature float32) Builder
}

type Builder struct {
    BuilderInterface
    thermostat Thermostat
}

func (builder Builder) SetId(id uint64) Builder {
    builder.thermostat.ID = id

    return builder
}

func (builder Builder) SetName(name string) Builder {
    builder.thermostat.Name = name

    return builder
}

func (builder Builder) SetAddress(address string) Builder {
    builder.thermostat.Address = address

    return builder
}

func (builder Builder) SetPort(port int) Builder {
    builder.thermostat.Port = port

    return builder
}

func (builder Builder) SetTypeArduino() Builder {
    builder.thermostat.Type = device.Arduino

    return builder
}

func (builder Builder) SetTypeRaspberryPi() Builder {
    builder.thermostat.Type = device.RaspberryPi

    return builder
}

func (builder Builder) SetAuto(auto bool) Builder {
    builder.thermostat.Auto = auto

    return builder
}

func (builder Builder) SetMinTemperature(temperature float32) Builder {
    builder.thermostat.MinTemperature = temperature

    return builder
}

func (builder Builder) Build() (Thermostat, error) {
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

func New() Builder {
    return Builder{}
}
