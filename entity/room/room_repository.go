package room

import (
    "errors"
    "github.com/avegao/iot-temp-service/util"
)

type RepositoryInterface interface {
    FindOneById(uint64 uint64) (Room, error)
    FindAll() ([]Room, error)
    Insert(thermostat Room) (Room, error)
    Update(thermostat Room) (Room, error)
    Delete(thermostat Room) (error)
}

type Repository struct {
    RepositoryInterface
}

func (repository Repository) FindOneById(id uint64) (Room, error) {
    query := "SELECT id, name, created_at, updated_at FROM rooms WHERE id = $1 AND deleted_at IS NULL"

    util.LogQuery(query, map[string]interface{}{"id": id})

    row := util.Pgsql.QueryRow(query, id)
    room := new(Room)
    err := row.Scan(
        &room.ID,
        &room.Name,
        &room.CreatedAt,
        &room.UpdatedAt)

    return *room, err
}

func (repository Repository) FindAll() ([]Room, error) {
    query := "SELECT id, name, created_at, updated_at FROM rooms WHERE deleted_at IS NULL"

    util.LogQuery(query, nil)

    rows, err := util.Pgsql.Query(query)

    if nil != err {
        return *new([]Room), err
    }

    defer rows.Close()

    zones := make([]Room, 0)

    for rows.Next() {
        room := new(Room)

        err = rows.Scan(
            &room.ID,
            &room.Name,
            &room.CreatedAt,
            &room.UpdatedAt)

        if nil != err {
            return *new([]Room), err
        }

        zones = append(zones, *room)
    }

    err = rows.Err()

    if nil != err {
        return *new([]Room), err
    }

    return zones, nil
}

func (repository Repository) Insert(thermostat Room) (Room, error) {
    return *new(Room), errors.New("not yet implemented")
}

func (repository Repository) Update(thermostat Room) (Room, error) {
    return *new(Room), errors.New("not yet implemented")
}

func (repository Repository) Delete(thermostat Room) (Room, error) {
    return *new(Room), errors.New("not yet implemented")
}
