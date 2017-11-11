package device

import (
    "github.com/avegao/iot-temp-service/util"
    "github.com/avegao/iot-temp-service/entity/zone"
    "github.com/avegao/iot-temp-service/entity/room"
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
        return "OTHER"
    }
}

const (
    Arduino     Type = 1
    RaspberryPi Type = 2
)

type Device struct {
    util.SoftDeletableEntity
    Name    string
    Address string
    Port    int
    Type    Type
    Zone    zone.Zone
    Room    room.Room
}

func (device Device) GetZone() (zone.Zone, error) {
    return *new(zone.Zone), errors.New("not yet implemented")
}

func (device Device) GetRoom() (room.Room, error) {
    return *new(room.Room), errors.New("not yet implemented")
}

