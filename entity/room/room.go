package room

import (
    "github.com/avegao/iot-temp-service/entity/zone"
    pb "github.com/avegao/iot-temp-service/resource/grpc/iot_temp"
    "github.com/avegao/iot-temp-service/util"
)

const (
    packageLogTag = "room."
    structLogTag  = packageLogTag + "room."
)

type Room struct {
    util.SoftDeletableEntity
    Name string
}

func (room Room) String() string {
    return room.Name
}

func (room Room) GetZone() (*zone.Zone, error) {
    return new(zone.Repository).FindOneByRoomId(room.ID)
}

func (room Room) ToGrpcResponse() (*pb.Room, error) {
    const logTag = structLogTag + "ToGrpcResponse"
    logger := util.GetContainer().GetLogger()
    logger.Debugf("%s - START", logTag)

    zoneObject, err := room.GetZone()

    if nil != err {
        logger.WithError(err).Debugf("%s - STOP -> With zone query", logTag)

        return nil, err
    }

    response := &pb.Room{
        Id:        room.ID,
        Name:      room.Name,
        Zone:      zoneObject.ToGrpcResponse(),
        CreatedAt: util.TimeToGrpcTimestamp(room.CreatedAt),
        UpdatedAt: util.TimeToGrpcTimestamp(room.UpdatedAt),
        DeletedAt: util.TimeToGrpcTimestamp(room.DeletedAt.Time),
    }

    logger.Debugf("%s - START", logTag)

    return response, nil
}
