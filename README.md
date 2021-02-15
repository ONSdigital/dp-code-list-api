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

The endpoint `/health` checks the connection to the database and returns one of:

- success (200, JSON "status":"OK")
- warning (429, JSON "status":"WARNING")
- failure (500, JSON "status":"CRITICAL")

### Configuration

| Environment variable         | Default                                | Description
| ---------------------------- | ---------------------------------------| -----------
| BIND_ADDR                    | :22400                                 | The host and port to bind to
| CODE_LIST_API_URL            | http://localhost:22400                 | The base URL for the code list API
| DATASET_API_URL              | http://localhost:22000                 | The base URL for the dataset API
| GRACEFUL_SHUTDOWN_TIMEOUT    | 5s                                     | The graceful shutdown timeout in seconds
| HEALTHCHECK_INTERVAL         | 30s                                    | Time between calls to healthchecks
| HEALTHCHECK_CRITICAL_TIMEOUT | 90s                                    | Timeout to consider a failing healthcheck critical
| DEFAULT_MAXIMUM_LIMIT        | 1000                                   | Default maximum limit for pagination
| DEFAULT_LIMIT                | 20                                     | Default limit for pagination
| DEFAULT_OFFSET               | 0                                      | Default offset for pagination

### License

Copyright © 2016-2021, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details
