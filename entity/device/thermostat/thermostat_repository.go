package thermostat

import (
    "errors"
    "github.com/avegao/iot-temp-service/util"
)

type RepositoryInterface interface {
    FindOneById(uint64 uint64) (Thermostat, error)
    FindAll() ([]Thermostat, error)
    Insert(thermostat Thermostat) (Thermostat, error)
    Update(thermostat Thermostat) (Thermostat, error)
    Delete(thermostat Thermostat) (error)
}

type Repository struct {
    RepositoryInterface
}

func (repository Repository) FindOneById(id uint64) (Thermostat, error) {
    query :=
        "SELECT " +
            "id, " +
            "name, " +
            "address, " +
            "port, " +
            "type, " +
            "auto, " +
            "min_temperature, " +
            "created_at, " +
            "updated_at, " +
            "id_zone, " +
            "id_room " +
            "FROM devices_thermostats " +
            "WHERE id = $1 " +
            "AND deleted_at IS NULL"

    util.LogQuery(query, map[string]interface{}{"id": id})

    database, err := util.GetContainer().GetDefaultDatabase()

    if nil != err {
        return Thermostat{}, err
    }

    row := database.QueryRow(query, id)
    thermostat := new(Thermostat)
    err = row.Scan(
        &thermostat.ID,
        &thermostat.Name,
        &thermostat.Address,
        &thermostat.Port,
        &thermostat.Type,
        &thermostat.Auto,
        &thermostat.MinTemperature,
        &thermostat.CreatedAt,
        &thermostat.UpdatedAt,
        &thermostat.ZoneID,
        &thermostat.RoomID,
    )

    return *thermostat, err
}

func (repository Repository) FindAll() ([]Thermostat, error) {
    query :=
        "SELECT " +
            "id, " +
            "name, " +
            "address, " +
            "port, " +
            "type, " +
            "auto, " +
            "min_temperature, " +
            "created_at, " +
            "updated_at, " +
            "id_zone, " +
            "id_room " +
            "FROM devices_thermostats " +
            "WHERE deleted_at IS NULL"

    util.LogQuery(query, nil)

    database, err := util.GetContainer().GetDefaultDatabase()

    if nil != err {
        return *new([]Thermostat), err
    }

    rows, err := database.Query(query)

    if nil != err {
        return *new([]Thermostat), err
    }

    defer rows.Close()

    thermostats := make([]Thermostat, 0)

    for rows.Next() {
        thermostat := new(Thermostat)
        err = rows.Scan(
            &thermostat.ID,
            &thermostat.Name,
            &thermostat.Address,
            &thermostat.Port,
            &thermostat.Type,
            &thermostat.Auto,
            &thermostat.MinTemperature,
            &thermostat.CreatedAt,
            &thermostat.UpdatedAt,
            &thermostat.ZoneID,
            &thermostat.RoomID,
        )

        thermostats = append(thermostats, *thermostat)
    }

    err = rows.Err()

    if nil != err {
        return *new([]Thermostat), err
    }

    return thermostats, nil
}

func (repository Repository) Insert(thermostat Thermostat) (Thermostat, error) {
    return *new(Thermostat), errors.New("not yet implemented")
}

func (repository Repository) Update(thermostat Thermostat) (Thermostat, error) {
    return *new(Thermostat), errors.New("not yet implemented")
}

func (repository Repository) Delete(thermostat Thermostat) (Thermostat, error) {
    return *new(Thermostat), errors.New("not yet implemented")
}
