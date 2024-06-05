--liquibase formatted sql

--changeset ammar.vepari:1 context:dev splitStatements:false
-- create schemas

CREATE SCHEMA IF NOT EXISTS employee;
