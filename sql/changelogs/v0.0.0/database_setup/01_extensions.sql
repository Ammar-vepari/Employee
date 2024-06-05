--liquibase formatted sql

--changeset ammar.vepari:1 context:dev splitStatements:true
-- create extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS pgcrypto;
