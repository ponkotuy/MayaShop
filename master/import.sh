#!/bin/sh

DB_NAME="maya"

find . -name "*.json" | xargs basename | cut -d. -f1 | while read i
do
    echo $i
    echo "db.$i.drop()" | arangosh --server.database $DB_NAME
    arangoimp --file "$i.json" --type json --collection $i --create-collection true --server.database $DB_NAME
done
