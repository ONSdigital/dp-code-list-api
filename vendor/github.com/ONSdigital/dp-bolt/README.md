# dp-bolt
Library adding a layer of abstraction around the low level Neo4j bolt driver handling the boilerplate Neo4j query code
 for you. Simply provide a query, parameters and a closure to extract the row data.
 
## Getting started
`go get github.com/ONSdigital/dp-bolt`

### Creating a bolt.DB
```go
pool, err := neo4j.NewClosableDriverPool("$bolt_url$", 1)
if err != nil {
	// handle error
}

db := bolt.New(pool)
defer db.Close()
```
### Querying for a single result
```
err = db.QueryForResult("MATCH (n) RETURN count(*)", nil, rowExtractor)
if err != nil {
    // handle error
}
```

### ResultExtractor closure
The final parameter of `db.QueryForResult()` is a `ResultExtractor` - a closure that enables you to customize how to 
handle the row data returned by your query (the underlying library returns row data as `[]interface{}` the closure you 
provide should handle casting the data to the expected type, handling any casting error and assigning to a variable for later use.
```go
type ResultExtractor func(r *Result) error
```
`Result` is a wrapper around the values returned by the underlying library - row `Data`,`Metadata` and `Index` which 
should provide everything you need to extract your response.  

#### Full example
```go
pool, err := neo4j.NewClosableDriverPool("$bolt_url$", 1)
if err != nil {
	// handle error
}

db := bolt.New(pool)
defer db.Close()

var count int64
rowExtractor := func(r *bolt.Result) error {
    var ok bool
    count, ok = r.Data[0].(int64)
    if !ok {
        return errors.New("failed to cast result to int64")
    }
    return nil
}

err = db.QueryForResult("MATCH (n) RETURN count(*)", nil, rowExtractor)
if err != nil {
    // handle error
}
// do something with count ...
```


