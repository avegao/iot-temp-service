package util

import "time"

type Entity struct {
    ID uint64
    CreatedAt time.Time
    UpdatedAt time.Time
}

type SoftDeletableEntity struct {
    Entity
    DeletedAt time.Time
}
