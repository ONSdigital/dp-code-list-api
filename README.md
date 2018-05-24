dp-code-list-api
================

An API used to navigate code lists

### Getting started

* Install mongodb
* Run `./scripts/setup.sh` to setup the code data

### Updating codes / code lists

Currently only a developer can add / update codes. To do this the following files need to be updated
* `scripts/codelists.json` (For all code lists)
* `scripts/codes.json` (For all codes)
* Run `./scripts/setup.sh` (This will drop the current data stored in mongodb)

### Configuration

| Environment variable         | Default                               | Description
| -------------------------    | ------------------------------------- | -----------
| BIND_ADDR                    | :22400                                | The host and port to bind to
| MONGODB_ADDR                 | localhost:27017                       | A URL to mongodb
| GRACEFUL_SHUTDOWN_TIMEOUT    | 5 seconds                             | The amount of time given when shutting down the API
| MONGODB_ADDR                 | localhost:27017                       | Connect address for MongoDB
| MONGODB_DATABASE             | codelists                             | Database name for collections for MongoDB
| MONGODB_CODELISTS_COLLECTION | codelists                             | Codelists collection name in MongoDB
| MONGODB_CODES_COLLECTION     | codes                                 | Codes colletion name in MongoDB

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2016-2017, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details
