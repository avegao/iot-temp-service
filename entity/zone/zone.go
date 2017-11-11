package zone

import "github.com/avegao/iot-temp-service/util"

type Zone struct {
    util.SoftDeletableEntity
    Name string
}

func (zone Zone) String() string {
    return zone.Name
}
