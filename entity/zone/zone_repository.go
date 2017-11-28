package zone

import (
    "database/sql"
    "errors"
    "github.com/avegao/gocondi"
    "github.com/avegao/iot-temp-service/util"
)

const structLogTag = packageLogTag + "repository."

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
}

func (repository Repository) FindOneById(id uint64) (*Zone, error) {
    const logTag = structLogTag + "FindOneById"
    logger := gocondi.GetContainer().GetLogger()
    logger.Debugf("%s - START", logTag)

    query := "SELECT id, name, created_at, updated_at FROM zones WHERE id = $1 AND deleted_at IS NULL"

    util.LogQuery(query, map[string]interface{}{"id": id})

    database, err := gocondi.GetContainer().GetDefaultDatabase()

    if nil != err {
        logger.Debugf("%s - STOP", logTag)

        if err == sql.ErrNoRows {
            logger.Debugf("%s - Zone not found", logTag)
        } else {
            logger.WithError(err).Error(logTag)

            err = nil
        }

        return nil, err
    }

    row := database.QueryRow(query, id)
    zone := new(Zone)
    err = row.Scan(
        &zone.ID,
        &zone.Name,
        &zone.CreatedAt,
        &zone.UpdatedAt)

    logger.Debugf("%s - END", logTag)

    return zone, err
}

func (repository Repository) FindAll() (*[]Zone, error) {
    query := "SELECT id, name, created_at, updated_at FROM zones WHERE deleted_at IS NULL"

    util.LogQuery(query, nil)

    database, err := gocondi.GetContainer().GetDefaultDatabase()

    if nil != err {
        return new([]Zone), err
    }

    rows, err := database.Query(query)

    if nil != err {
        return new([]Zone), err
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
            return new([]Zone), err
        }

        zones = append(zones, *zone)
    }

    err = rows.Err()

    if nil != err {
        return new([]Zone), err
    }

    return &zones, nil
}

func (repository Repository) FindOneByRoomId(id uint64) (*Zone, error) {
    const logTag = structLogTag + "FindOneById"
    logger := gocondi.GetContainer().GetLogger()
    logger.Debugf("%s - START", logTag)

    query := "SELECT id, name, created_at, updated_at FROM zones WHERE id IN (SELECT id_zone FROM rooms WHERE id = $1) AND deleted_at IS NULL"

    util.LogQuery(query, map[string]interface{}{"id": id})

    database, err := gocondi.GetContainer().GetDefaultDatabase()

    if nil != err {
        logger.WithError(err).Debugf("%s - STOP -> Error with database", logTag)

        return nil, err
    }

    row := database.QueryRow(query, id)
    zone := new(Zone)
    err = row.Scan(
        &zone.ID,
        &zone.Name,
        &zone.CreatedAt,
        &zone.UpdatedAt)

    if nil != err {
        logger.WithError(err).Debugf("%s - STOP -> Error reading row", logTag)

        return nil, err
    }

    logger.Debugf("%s - END", logTag)

    return zone, nil
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
