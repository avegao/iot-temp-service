package room

import (
    "github.com/avegao/iot-temp-service/entity/zone"
    "github.com/avegao/iot-temp-service/util"
)

type Room struct {
    util.SoftDeletableEntity
    Name string
}

func (room Room) String() string {
    return room.Name
}

func (room Room) GetZone() (zone.Zone, error) {

    return new(zone.Repository).FindOneByRoomId(room.ID)
}
