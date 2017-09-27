#!/usr/bin/env bash

mongo mongodb://localhost:27017/codelists <<EOF
db.dropDatabase();
db.codes.ensureIndex({"code":1,"links.code_list.id":1},{"background":true});
EOF

scriptDir=$( dirname "${BASH_SOURCE[0]}" )
mongoimport --db codelists --collection codelists --file $scriptDir/codelists.json
mongoimport --db codelists --collection codes --file $scriptDir/codes.json
