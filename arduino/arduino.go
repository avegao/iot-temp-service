package arduino

import (
    "github.com/avegao/gocondi"
    "google.golang.org/grpc"
    "github.com/avegao/iot-temp-service/resource/grpc"
    "golang.org/x/net/context"
    "github.com/avegao/iot-temp-service/util"
)

var (
    connection *grpc.ClientConn
    client     arduino_service.ArduinoClient
    address    string
)

func createConnection() {
    var grpcOptions []grpc.DialOption

    //if *tls {
    //    if *caFile == "" {
    //        *caFile = testdata.Path("ca.pem")
    //    }
    //    grpcCredentials, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
    //    if err != nil {
    //        log.Fatalf("Failed to create TLS credentials %v", err)
    //    }
    //    opts = append(opts, grpc.WithTransportCredentials(grpcCredentials))
    //} else {
    //    opts = append(opts, grpc.WithInsecure())
    //}

    grpcOptions = append(grpcOptions, grpc.WithInsecure())
    address = util.FromGeneric(gocondi.GetContainer().GetParameter("iot_arduino_temp_server_address"))
    newConnection, err := grpc.Dial(address, grpcOptions...)

    if nil != err {
        gocondi.GetContainer().GetLogger().WithError(err).Fatalf("Fail to connect with %s", address)
    }

    connection = newConnection

    gocondi.GetContainer().GetLogger().Debugf("gRPC connection status with %v = %s", address, connection.GetState().String())
}

func CloseConnection() {
    if nil != connection {
        connection.Close()
    }
}

func createClient() {
    if nil == connection {
        createConnection()
    }

    if nil == connection {
        gocondi.GetContainer().GetLogger().Panic("connection null")
    }

    client = arduino_service.NewArduinoClient(connection)
}

func checkClientStatus() {
    if nil == client {
        createClient()
    }

    gocondi.GetContainer().GetLogger().Debugf("gRPC connection status with %v = %s", address, connection.GetState())
}

func GetTemperature(request arduino_service.ArduinoRequest) (float32, error) {
    checkClientStatus()

    response, err := client.GetTemperature(context.Background(), &request)

    if nil != err {
        return float32(0), err
    }

    return response.Temperature, nil
}

func IsPower(request arduino_service.ArduinoRequest) (bool, error) {
    checkClientStatus()

    response, err := client.IsPower(context.Background(), &request)

    return response.Power, err
}

func PowerOff(request arduino_service.ArduinoRequest) (bool, error) {
    checkClientStatus()

    response, err := client.PowerOff(context.Background(), &request)

    return response.Power, err
}

func PowerOn(request arduino_service.ArduinoRequest) (bool, error) {
    checkClientStatus()

    response, err := client.PowerOn(context.Background(), &request)

    return response.Power, err
}
