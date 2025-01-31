#!/bin/bash
source ./.env
BACKUP_DIR="./backups"
LOG_FILE="$BACKUP_DIR/backup.log"
DATE=$(date +"%Y-%m-%d_%H-%M-%S")
BACKUP_FILE="$BACKUP_DIR/backup_$DATE.sql"
mkdir -p "$BACKUP_DIR"
export PGPASSWORD="$DB_PASSWORD"
pg_dump --username="$DB_USER" --host="$DB_HOST" --port="$DB_PORT" --dbname="$DB_NAME" --format=plain --file="$BACKUP_FILE"
EXIT_CODE=$?
unset PGPASSWORD
if [ $EXIT_CODE -eq 0 ]; then
    echo "[$(date +"%Y-%m-%d %H:%M:%S")] Backup PostgreSQL réussi : $BACKUP_FILE" | tee -a "$LOG_FILE"
else
    echo "[$(date +"%Y-%m-%d %H:%M:%S")] ERREUR: La sauvegarde a échoué !" | tee -a "$LOG_FILE"
    exit 1
fi
