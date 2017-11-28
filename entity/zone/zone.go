package zone

import (
    pb "github.com/avegao/iot-temp-service/resource/grpc/iot_temp"
    "github.com/avegao/iot-temp-service/util"
)

const packageLogTag = "zone."

type Zone struct {
    util.SoftDeletableEntity
    Name string
}

func (zone Zone) String() string {
    return zone.Name
}

func (zone Zone) ToGrpcResponse() *pb.Zone {
    return &pb.Zone{
        Id:        zone.ID,
        Name:      zone.Name,
        CreatedAt: util.TimeToGrpcTimestamp(zone.CreatedAt),
        UpdatedAt: util.TimeToGrpcTimestamp(zone.UpdatedAt),
        DeletedAt: util.TimeToGrpcTimestamp(zone.DeletedAt.Time),
    }
}
