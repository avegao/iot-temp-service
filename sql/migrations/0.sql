CREATE OR REPLACE FUNCTION update_updated_at_column()
  RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

CREATE TABLE "public"."zones" (
  "id"         SERIAL8,
  "name"       VARCHAR(255)   NOT NULL,
  "created_at" TIMESTAMPTZ(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ(0) NOT NULL DEFAULT NOW(),
  "deleted_at" TIMESTAMPTZ(0),
  CONSTRAINT pk_zones PRIMARY KEY ("id")
);

CREATE TRIGGER "tg_zones_updated_at"
  BEFORE UPDATE
  ON "public"."zones"
  FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TABLE "public"."rooms" (
  "id"         SERIAL8,
  "name"       VARCHAR(255)   NOT NULL,
  "created_at" TIMESTAMPTZ(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ(0) NOT NULL DEFAULT NOW(),
  "deleted_at" TIMESTAMPTZ(0),
  CONSTRAINT pk_rooms PRIMARY KEY ("id")
);

CREATE TRIGGER "tg_rooms_updated_at"
  BEFORE UPDATE
  ON "public"."rooms"
  FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TABLE "public"."devices" (
  "id"         SERIAL8,
  "name"       VARCHAR(255)   NOT NULL,
  "address"    VARCHAR(255)   NOT NULL,
  "port"       VARCHAR(10)    NOT NULL,
  "type"       INT2           NOT NULL DEFAULT 0,
  "id_zone"    INT8           NOT NULL,
  "id_room"    INT8           NOT NULL,
  "created_at" TIMESTAMPTZ(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ(0) NOT NULL DEFAULT NOW(),
  "deleted_at" TIMESTAMPTZ(0),
  CONSTRAINT pk_devices PRIMARY KEY ("id"),
  CONSTRAINT fk_devices_zones FOREIGN KEY ("id_zone") REFERENCES "public"."zones" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT fk_devices_rooms FOREIGN KEY ("id_room") REFERENCES "public"."rooms" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TRIGGER "tg_devices_updated_at"
  BEFORE UPDATE
  ON "public"."devices"
  FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TABLE "public"."devices_thermostats" (
  "auto" BOOLEAN NOT NULL DEFAULT FALSE,
  "min_temperature" NUMERIC(4, 2) NULL DEFAULT NULL,
  CONSTRAINT pk_devices_thermostats PRIMARY KEY ("id")
) INHERITS ("public"."devices");

CREATE TRIGGER "tg_devices_thermostats_updated_at"
  BEFORE UPDATE
  ON "public"."devices_thermostats"
  FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
