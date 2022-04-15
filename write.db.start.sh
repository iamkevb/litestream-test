#!/bin/bash
mkdir db
echo Creating write database
echo Existing tables:
sqlite3 db/write.db ".tables"

echo
echo Starting litestream on port 9090
LITESTREAM_HOST=127.0.0.1 LITESTREAM_PORT=9090 ./litestream replicate -config write.db.yml