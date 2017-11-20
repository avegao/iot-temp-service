package zone

import (
    "errors"
    "github.com/avegao/iot-temp-service/util"
    "database/sql"
)

type repositoryInterface interface {
    Delete(thermostat Zone) (error)
    FindAll() ([]Zone, error)
    FindOneById(uint64 uint64) (Zone, error)
    FindOneByRoomId(id uint64) (Zone, error)
    Insert(thermostat Zone) (Zone, error)
    Update(thermostat Zone) (Zone, error)
}

type Repository struct {
    repositoryInterface
    DB sql.DB
}

func (repository Repository) FindOneById(id uint64) (Zone, error) {
    query := "SELECT id, name, created_at, updated_at FROM zones WHERE id = $1 AND deleted_at IS NULL"

    util.LogQuery(query, map[string]interface{}{"id": id})

    row := repository.DB.QueryRow(query, id)
    zone := new(Zone)
    err := row.Scan(
        &zone.ID,
        &zone.Name,
        &zone.CreatedAt,
        &zone.UpdatedAt)

    return *zone, err
}

func (repository Repository) FindAll() ([]Zone, error) {
    query := "SELECT id, name, created_at, updated_at FROM zones WHERE deleted_at IS NULL"

    util.LogQuery(query, nil)

    rows, err := repository.DB.Query(query)

    if nil != err {
        return *new([]Zone), err
    }

    defer rows.Close()

    zones := make([]Zone, 0)

    for rows.Next() {
        zone := new(Zone)

        err = rows.Scan(
            &zone.ID,
            &zone.Name,
            &zone.CreatedAt,
            &zone.UpdatedAt)

        if nil != err {
            return *new([]Zone), err
        }

        zones = append(zones, *zone)
    }

    err = rows.Err()

    if nil != err {
        return *new([]Zone), err
    }

    return zones, nil
}

func (repository Repository) FindOneByRoomId(id uint64) (Zone, error) {
    query := "SELECT id, name, created_at, updated_at FROM zones WHERE id IN (SELECT id_zone FROM rooms WHERE id = $1) AND deleted_at IS NULL"

    util.LogQuery(query, map[string]interface{}{"id": id})

    row := repository.DB.QueryRow(query, id)
    zone := new(Zone)
    err := row.Scan(
        &zone.ID,
        &zone.Name,
        &zone.CreatedAt,
        &zone.UpdatedAt)

    return *zone, err
}

func (repository Repository) Insert(thermostat Zone) (Zone, error) {
    return *new(Zone), errors.New("not yet implemented")
}

func (repository Repository) Update(thermostat Zone) (Zone, error) {
    return *new(Zone), errors.New("not yet implemented")
}

func (repository Repository) Delete(thermostat Zone) (Zone, error) {
    return *new(Zone), errors.New("not yet implemented")
}
