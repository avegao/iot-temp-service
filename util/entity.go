package util

import (
    "github.com/avegao/gocondi"
    "time"
    "github.com/lib/pq"
    "github.com/sirupsen/logrus"
)

type Entity struct {
    ID        uint64
    CreatedAt time.Time
    UpdatedAt time.Time
}

type SoftDeletableEntity struct {
    Entity
    DeletedAt pq.NullTime
}

func LogQuery(query string, parameters map[string]interface{}) {
    gocondi.GetContainer().GetLogger().WithFields(logrus.Fields{"query": query, "parameters": parameters}).Debugf("Query executed")
}
