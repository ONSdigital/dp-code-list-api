#!/usr/bin/env bash

# setup.sh [ --env ]
#
# use `--env` if running against an environment

MONGODB_ADDR=${MONGODB_ADDR:-mongodb://localhost:27017}
auth_source=
use_env=

if [[ $1 == --env ]]; then
        use_env=1
        shift
fi

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
        if [[ -n $use_env ]]; then
                mongoimport --uri $MONGODB_ADDR/codelists$auth_source --collection $collection --file <(sed 's/localhost:22400/localhost:10500/g' $file) || exit 2
        else
                mongoimport --uri $MONGODB_ADDR/codelists$auth_source --collection $collection --file $file || exit 2
        fi
}

import_to codelists codelists.json
import_to codes codes/time.json
import_to codes codes/cpi1dim1aggid.json
import_to codes codes/cpih1dim1aggid.json
import_to codes codes/uk-only.json
import_to codes codes/midyearpopgeography.json
import_to codes codes/midyearpopage.json
import_to codes codes/midyearpopsex.json
