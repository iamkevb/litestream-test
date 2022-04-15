1. Run `write.db.start.sh`
   1. Creates `db/write.db`
   1. Starts litestream (from binary root) replicated to `db/backup`
1. Run `read.db.start.sh`
   1. Starts litestream read replica, creating `db/read.db`
1. Run `go run driver.go -create <tablename>`
   1. Table is created in write db
   1. Read db is queried until results are found (new table has been replicated)
