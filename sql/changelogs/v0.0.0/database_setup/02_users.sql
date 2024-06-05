--liquibase formatted sql

--changeset ammar.vepari:1 context:dev splitStatements:false 
-- create user
DO
$$
BEGIN
    IF NOT EXISTS (SELECT * FROM pg_user where usename = '${DB_USER}' ) THEN
        create user ${DB_USER} with encrypted password '${DB_PASS}';
    END IF;
END $$;

GRANT USAGE ON SCHEMA employee to ${DB_USER};

GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA employee to ${DB_USER};
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA employee to ${DB_USER};

ALTER DEFAULT PRIVILEGES IN SCHEMA employee GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO ${DB_USER};
ALTER DEFAULT PRIVILEGES IN SCHEMA employee GRANT ALL PRIVILEGES ON SEQUENCES TO ${DB_USER};
