#!/usr/bin/env bash

mongo mongodb://localhost:27017/codelists <<EOF
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
        mongoimport --db codelists --collection $collection --file $file || exit 2
}

import_to codelists codelists.json
import_to codes codes/time.json
import_to codes codes/cpi1dim1aggid.json
import_to codes codes/54ff5089-ea78-45ef-afa2-0dfe58f89497.json
import_to codes codes/4e1a21a9-3fb9-4a71-b2ad-9be29457236b.json
