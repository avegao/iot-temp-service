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
            "updated_at " +
            "FROM devices_thermostats " +
            "WHERE id = $1 " +
            "AND deleted_at IS NULL"

    util.LogQuery(query, map[string]interface{}{"id": id})

    row := util.Pgsql.QueryRow(query, id)
    thermostat := new(Thermostat)
    err := row.Scan(
        &thermostat.ID,
        &thermostat.Name,
        &thermostat.Address,
        &thermostat.Port,
        &thermostat.Type,
        &thermostat.Auto,
        &thermostat.MinTemperature,
        &thermostat.CreatedAt,
        &thermostat.UpdatedAt)

    return *thermostat, err
}

func (repository Repository) FindAll() ([]Thermostat, error) {
    return *new([]Thermostat), errors.New("not yet implemented")
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
