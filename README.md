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

Run `make debug` (or `make acceptance` when running acceptance tests)

### Healthcheck

The endpoint `/healthcheck` checks the connection to the database and returns one of:

- success (200, JSON "status":"OK")
- failure (500, JSON "status":"error").

### Configuration

| Environment variable        | Default                                | Description
| --------------------------- | ---------------------------------------| -----------
| BIND_ADDR                   | :22000                                 | The host and port to bind to
| NEO4J_BIND_ADDRESS          | bolt://localhost:7687                  | The address of the neo4j database to retrieve dataset data from
| NEO4J_POOL_SIZE             | 5                                      | The number of neo4j connections to pool
| NEO4J_CODE_LIST_LABEL       | code_list                              | The code list node label identifier in neo4j
| GRACEFUL_SHUTDOWN_TIMEOUT   | 5s                                     | The graceful shutdown timeout in seconds


### License

Copyright © 2016-2017, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details
