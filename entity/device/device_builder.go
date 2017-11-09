package device

import (
    "github.com/sirupsen/logrus"
    "github.com/avegao/iot-temp-service/util"
    "errors"
)

type deviceBuilderInterface interface {
    util.Builder
    SetId(id uint64) deviceBuilder
    SetName(name string) deviceBuilder
    SetAddress(address string) deviceBuilder
    SetPort(port int) deviceBuilder
}

type deviceBuilder struct {
    deviceBuilderInterface
    device Device
}

func (deviceBuilder deviceBuilder) SetId(id uint64) deviceBuilder {
    deviceBuilder.device.ID = id

    return deviceBuilder
}

func (deviceBuilder deviceBuilder) SetName(name string) deviceBuilder {
    deviceBuilder.device.Name = name

    return deviceBuilder
}

func (deviceBuilder deviceBuilder) SetAddress(address string) deviceBuilder {
    deviceBuilder.device.Address = address

    return deviceBuilder
}

func (deviceBuilder deviceBuilder) SetPort(port int) deviceBuilder {
    deviceBuilder.device.Port = port

    return deviceBuilder
}

func (deviceBuilder deviceBuilder) SetTypeArduino() deviceBuilder {
    deviceBuilder.device.Type = Arduino

    return deviceBuilder
}

func (deviceBuilder deviceBuilder) SetTypeRaspberryPi() deviceBuilder {
    deviceBuilder.device.Type = RaspberryPi

    return deviceBuilder
}

func (deviceBuilder deviceBuilder) Build() (Device, error) {
    if deviceBuilder.device.ID < 0 {
        logrus.Warnf("ID must be >= 0. Setting to 0. This result on a new device.")

        deviceBuilder.device.ID = 0
    }

    if util.Empty(deviceBuilder.device.Name) {
        return *new(Device), errors.New("device name is required")
    }

    if util.Empty(deviceBuilder.device.Address) {
        return *new(Device), errors.New("device address is required")
    }

    if 0 == deviceBuilder.device.Port {
        return *new(Device), errors.New("device port is required")
    }

    if 0 == deviceBuilder.device.Type {
        return *new(Device), errors.New("device type is required")
    } else if deviceBuilder.device.Type != Arduino {
        logrus.Panicf("%s not support yet", deviceBuilder.device.Type.String())
    }

    return deviceBuilder.device, nil
}

func New() deviceBuilder {
    return deviceBuilder{}
}

