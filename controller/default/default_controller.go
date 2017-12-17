package default_controller

import (
    "database/sql"
    "github.com/avegao/gocondi"
    "github.com/avegao/iot-temp-service/entity/device/thermostat"
    pb "github.com/avegao/iot-temp-service/resource/grpc/iot_temp"
    "github.com/golang/protobuf/ptypes/empty"
    "golang.org/x/net/context"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

const (
    packageLogTag = "default_controller."
    structLogTag  = packageLogTag + "controller."
)

type Controller struct {
    pb.ThermostatServiceServer
}

func (controller Controller) FindAll(ctx context.Context, in *empty.Empty) (*pb.ThermostatArray, error) {
    const logTag = structLogTag + "FindAll()"
    logger := gocondi.GetContainer().GetLogger()
    logger.Debugf("%s - START", logTag)

    repository := new(thermostat.Repository)
    thermostats, err := repository.FindAll()

    if nil != err {
        logger.WithError(err).Error(logTag)
        logger.Debugf("%s - STOP", logTag)

        return nil, status.Error(codes.Internal, err.Error())
    }

    response, err := thermostatsToArrayResponse(thermostats)

    if nil != err {
        logger.WithError(err).Error(logTag)
        logger.Debugf("%s - STOP", logTag)

        return nil, err
    }

    logger.Debugf("%s - END", logTag)

    return response, nil
}

func (controller Controller) FindOneById(ctx context.Context, in *pb.ByIdReqest) (*pb.Thermostat, error) {
    const logTag = structLogTag + "FindOneById()"
    logger := gocondi.GetContainer().GetLogger()
    logger.Debugf("%s - START", logTag)

    repository := new(thermostat.Repository)
    thermostatObject, err := repository.FindOneById(in.Id)

    if nil != err {
        logger.WithError(err).Error(logTag)
        logger.Debugf("%s - STOP", logTag)

        return nil, status.Error(codes.Internal, err.Error())
    }

    response, err := thermostatObject.ToGrpcResponse()

    if nil != err {
        logger.WithError(err).Error(logTag)
        logger.Debugf("%s - STOP", logTag)

        return nil, status.Error(codes.Internal, err.Error())
    }

    logger.Debugf("%s - END", logTag)

    return response, nil
}

func (controller Controller) GetTemperatureById(ctx context.Context, in *pb.ByIdReqest) (*pb.GetTemperatureResponse, error) {
    const logTag = structLogTag + "GetTemperatureById()"
    logger := gocondi.GetContainer().GetLogger()
    logger.Debugf("%s - START", logTag)

    repository := new(thermostat.Repository)
    thermostatObject, err := repository.FindOneById(in.Id)

    if nil != err {
        var code codes.Code

        if nil != sql.ErrNoRows {
            logger.Debugf("%s - STOP -> Error ", logTag)

            code = codes.NotFound
        } else {
            logger.WithError(err).Error(logTag)
            logger.WithError(err).Debugf("%s - STOP -> Error ", logTag)

            code = codes.Internal
        }

        return nil, status.Error(code, err.Error())
    }

    temperature, err := thermostatObject.GetTemperature()

    if nil != err {
        logger.WithError(err).Error(logTag)
        logger.WithError(err).Debugf("%s - STOP -> Error ", logTag)

        return nil, status.Error(codes.Internal, err.Error())
    }

    response := &pb.GetTemperatureResponse{Temperature: temperature}

    logger.Debugf("%s - END", logTag)

    return response, nil
}

func thermostatsToArrayResponse(thermostats []thermostat.Thermostat) (*pb.ThermostatArray, error) {
    const logTag = packageLogTag + "thermostatsToArrayResponse()"
    logger := gocondi.GetContainer().GetLogger()
    logger.Debugf("%s - START", logTag)

    responseArray := make([]*pb.Thermostat, 0)

    for _, thermostatObject := range thermostats {
        response, err := thermostatObject.ToGrpcResponse()

        if nil != err {
            logger.WithError(err).Error(logTag)
            logger.Debugf("%s - STOP", logTag)

            code := codes.Internal

            if err == sql.ErrNoRows {
                code = codes.NotFound
            }

            return nil, status.Error(code, err.Error())
        }

        responseArray = append(responseArray, response)
    }

    logger.Debugf("%s - END", logTag)

    return &pb.ThermostatArray{Thermostats: responseArray}, nil
}
