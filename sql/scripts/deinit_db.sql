SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE pid <> pg_backend_pid() AND datname = '${DB_NAME}';

SELECT 'DROP DATABASE ${DB_NAME};'
WHERE EXISTS (SELECT FROM pg_database WHERE datname = '${DB_NAME}')\gexec
