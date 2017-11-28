package util

import (
    "github.com/golang/protobuf/ptypes/timestamp"
    "time"
)

func TimeToGrpcTimestamp(value time.Time) *timestamp.Timestamp {
    return &timestamp.Timestamp{Seconds: int64(value.Second())}
}
