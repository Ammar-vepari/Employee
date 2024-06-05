-- liquibase formatted sql
-- changeset ammar.vepari:1 context:dev splitStatements:false
CREATE TABLE IF NOT EXISTS employee.details (
  id UUID PRIMARY KEY,
  name TEXT,
  position TEXT ,
  salary NUMERIC(10, 2),
  created_by TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
  updated_by TEXT,
  updated_at TIMESTAMPTZ
);
