#!/usr/bin/env bash

# setup.sh [ --env ]
#
# use `--env` if running against an environment

MONGODB_ADDR=${MONGODB_ADDR:-mongodb://localhost:27017}
auth_source=
use_env=
file=
collection=

for var in "$@"
do
    if [[ $1 == --codelist ]]; then
        shift
        collection=codelist
    elif [[ $1 == --codes ]]; then
       shift
       collection=codes
    elif [[ $1 == --file ]]; then
        shift
        file=$1
        shift
        if [ ! -f $file ]; then
            echo "$1 does not exist"
            exit 1
        fi
    elif [[ $1 == --env ]]; then
        use_env=1
        shift
    fi
done

if [[ -z $collection ]]; then
    echo "must include --codelist or --codes"
    exit 1
fi

if [[ $MONGODB_ADDR == *@* ]]; then
        auth_source="?authSource=admin"
fi

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

import_to $collection $file
