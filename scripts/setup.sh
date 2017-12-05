#!/usr/bin/env bash

MONGODB_ADDR=${MONGODB_ADDR:-mongodb://localhost:27017}
auth_source=
if [[ $MONGODB_ADDR == *@* ]]; then
        auth_source="?authSource=admin"
fi

mongo $MONGODB_ADDR/codelists$auth_source <<EOF
db.dropDatabase();
db.codelists.ensureIndex({"id":1},{"background":true});
db.codes.ensureIndex({"links.code_list.id":1},{"background":true});
db.codes.ensureIndex({"code":1,"links.code_list.id":1},{"background":true});
EOF

scriptDir=$( dirname "${BASH_SOURCE[0]}" )
cd $scriptDir || exit 2

import_to() {
        local collection=$1 file=$2; shift 2
        echo Importing $file ...
        mongoimport --uri $MONGODB_ADDR/codelists$auth_source --collection $collection --file $file || exit 2
}

import_to codelists codelists.json
import_to codes codes/time.json
import_to codes codes/cpi1dim1aggid.json
import_to codes codes/cpih1dim1aggid.json
import_to codes codes/uk-only.json
