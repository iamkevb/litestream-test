#!/bin/bash
mkdir db
echo Starting litestream read replicate of 127.0.0.1:9090
LITESTREAM_HOST=127.0.0.1 LITESTREAM_PORT=9090 WORKING_DIR=$PWD ./litestream replicate -config read.db.yml