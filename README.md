dp-code-list-api
================

An API used to navigate code lists

### Installation

#### Database

- Run `brew install neo4j`
- Configure neo4j, edit `/usr/local/Cellar/neo4j/*/libexec/conf/neo4j.conf`
- Set `dbms.security.auth_enabled=false`
- Run `brew services restart neo4j`

#### Getting started

1) Import the relevant code-list data using the [code list scripts](https://github.com/ONSdigital/dp-code-list-scripts).

2) Run `make debug` (or `make acceptance` when running acceptance tests)

### Healthcheck

The endpoint `/healthcheck` checks the connection to the database and returns one of:

- success (200, JSON "status":"OK")
- failure (500, JSON "status":"error").

### Configuration

| Environment variable          | Default                                | Description
| ----------------------------- | ---------------------------------------| -----------
| BIND_ADDR                     | :22400                                 | The host and port to bind to
| CODE_LIST_API_URL             | http://localhost:22400                 | The base URL for the code list API
| DATASET_API_URL               | http://localhost:22000                 | The base URL for the dataset API
| GRACEFUL_SHUTDOWN_TIMEOUT     | 5s                                     | The graceful shutdown timeout in seconds
| HEALTHCHECK_INTERVAL          | 30s                                    | Time between calls to healthchecks
| HEALTHCHECK_RECOVERY_INTERVAL | 5s                                     | Time between calls to healthchecks while failing


### License

Copyright © 2016-2019, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details
