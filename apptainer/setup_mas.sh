#!/bin/bash
set -xeu

#setup postgres
export LD_LIBRARY_PATH="${LD_LIBRARY_PATH-:}:/usr/local/lib"
export PGUSER=postgres
export PGDATA=/pg_data

mkdir $PGDATA
chown $PGUSER:$PGUSER $PGDATA
mkdir /run/postgresql && chown $PGUSER: /run/postgresql

su -c "initdb -A trust -U $PGUSER -D $PGDATA" -l $PGUSER
su -c "pg_ctl -D /pg_data -l logfile start " -l $PGUSER

# setup mas
ls /gsky/share/mas
(cd /gsky/share/mas && psql -f schema.sql)
(cd /gsky/share/mas && psql -f mas.sql)

su -c "pg_ctl -D /pg_data stop" -l $PGUSER