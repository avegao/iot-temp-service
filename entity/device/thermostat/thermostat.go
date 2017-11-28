package thermostat

import (
    "database/sql"
    "errors"
    "fmt"
    "github.com/avegao/iot-temp-service/arduino"
    "github.com/avegao/iot-temp-service/entity/device"
    "github.com/avegao/iot-temp-service/resource/grpc"
    pb "github.com/avegao/iot-temp-service/resource/grpc/iot_temp"
    "github.com/avegao/iot-temp-service/util"
    "github.com/sirupsen/logrus"
)

const (
    packageLogTag = "thermostat."
    structLogTag  = packageLogTag + "thermostat."
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

func (thermostat Thermostat) ToGrpcResponse() (*pb.Thermostat, error) {
    const logTag = structLogTag + "ToGrpcResponse"
    logger := util.GetContainer().GetLogger()
    logger.Debugf("%s - START", logTag)

    var roomResponse *pb.Room
    room, err := thermostat.GetRoom()

    if nil != err {
        if sql.ErrNoRows != err {
            logger.WithError(err).Debugf("%s - STOP -> Error with room query", logTag)

            return nil, err
        }
    } else {
        roomResponse, err = room.ToGrpcResponse()

        if nil != err {
            logger.WithError(err).Debugf("%s - STOP -> Error with room grpc response", logTag)

            return nil, err
        }
    }

    var zoneResponse *pb.Zone

    if (nil != roomResponse && nil == roomResponse.Zone) || nil == roomResponse {
        zone, err := thermostat.GetZone()

        if nil != err {
            logger.WithError(err).Debugf("%s - STOP -> Error with zone query", logTag)

            return nil, err
        }

        zoneResponse = zone.ToGrpcResponse()
    } else {
        zoneResponse = roomResponse.Zone
    }

    response := &pb.Thermostat{
        Id:        thermostat.ID,
        Name:      thermostat.Name,
        Address:   thermostat.Address,
        Port:      int32(thermostat.Port),
        Type:      pb.Thermostat_DeviceType(thermostat.Type),
        Zone:      zoneResponse,
        Room:      roomResponse,
        CreatedAt: util.TimeToGrpcTimestamp(thermostat.CreatedAt),
        UpdatedAt: util.TimeToGrpcTimestamp(thermostat.UpdatedAt),
        DeletedAt: util.TimeToGrpcTimestamp(thermostat.DeletedAt.Time),
    }

    logger.Debugf("%s - END", logTag)

    return response, nil
}
