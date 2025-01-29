#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE ROLE $PG_REP_USER WITH REPLICATION PASSWORD '$PG_REP_PASSWORD' LOGIN;
EOSQL

# Configuration du fichier pg_hba.conf pour autoriser la réplication
echo "host replication $PG_REP_USER 0.0.0.0/0 md5" >> "$PGDATA/pg_hba.conf"
