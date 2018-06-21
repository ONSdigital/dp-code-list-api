Migration from mongo to neo4j
================

Scripts to update the information stored in mongo ahead of the codelist/geography work.

Uses the API to generate a cypher query in the browser which can then be run against Neo to update the data.

### adddatatobepublished.html

Calculates dataset/edition and version adds these as labels to the instances in Neo4j

dataset: "string"
edition: "string"
version: integer

Also add published flag

published_flag : true

### codelists.html

Calulates the codelists and codes needed and writes the script tho import these and add the relationships between them.

Note: If run more than once will create duplicate entries
