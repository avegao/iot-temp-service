package device

import (
	"github.com/avegao/iot-temp-service/entity/room"
	"github.com/avegao/iot-temp-service/entity/zone"
	"github.com/avegao/iot-temp-service/util"
)

type Type int

func (deviceType Type) String() string {
	switch deviceType {
	case Arduino:
		return "ARDUINO"
	case RaspberryPi:
		return "RASPBERRY_PI"
	case ElGato:
		return "EL_GATO"
	case Xiaomi:
		return "XIAOMI"
	default:
		return "OTHER"
	}
}

const (
	Arduino     Type = 1
	RaspberryPi Type = 2
	ElGato      Type = 3
	Xiaomi      Type = 4
)

type Device struct {
	util.SoftDeletableEntity
	Name    string
	Address string
	Port    int
	Type    Type
	RoomID  uint64
}

func (device Device) GetZone() (*zone.Zone, error) {
	repository := new(zone.Repository)

	return repository.FindOneById(device.ZoneID)
}

func (device Device) GetRoom() (*room.Room, error) {
	repository := new(room.Repository)

	return repository.FindOneById(device.RoomID)
}
