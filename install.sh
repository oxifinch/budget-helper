#!/bin/bash

go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest

go get github.com/mattn/go-sqlite3
go get github.com/volatiletech/sqlboiler/v4
go get github.com/volatiletech/sqlboiler/v4/queries/qm
go get github.com/volatiletech/null/v8

sqlite3 app-db.db ".read init.sql"

sqlboiler sqlite3
