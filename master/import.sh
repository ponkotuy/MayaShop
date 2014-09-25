#!/bin/sh

DB_NAME="maya"

find . -name "*.json" | xargs -n1 basename | cut -d. -f1 | while read i
do
    echo $i
    jq -c "." $i.json > tmp         # arangoimpがcompactにしたJSONしか読んでくれない
    echo "db.$i.drop()" | arangosh --server.database $DB_NAME
    arangoimp --file tmp --type json --collection $i --create-collection true --server.database $DB_NAME
    rm tmp
done
