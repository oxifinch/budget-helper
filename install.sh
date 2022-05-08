#!/bin/bash

echo "This script will install project dependencies and remove the old database."

read -p "Continue? [Y/n] " input
case $input in
    "Y" | "y" | "")
        echo "Installing..."
        ;;
    *)
        echo "Aborted."
        exit 1
        ;;
esac

echo "- Dowloading dependencies."
go get gorm.io/gorm
go get gorm.io/driver/sqlite

echo "- Cleaning up module and updating."
go mod tidy
go get -u 

if [[ -f "./app-db.db" ]]; then
    echo "- Removing old database file."
    rm ./app-db.db
fi

echo "Done!"
