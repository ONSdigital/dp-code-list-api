package main

import (
	"fmt"
	neo4j "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"os"
	"github.com/ONSdigital/dp-bolt/bolt"
	"errors"
	"strconv"
	"log"
)

// Example of how to use library
func main() {
	pool, err := neo4j.NewClosableDriverPool("bolt://localhost:7687", 5)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	db := bolt.New(pool)
	defer func() {
		fmt.Println("closing db connection pool")
		db.Close()
	}()

	query := fmt.Sprintf("MATCH (cl:_code_list:`_name_%s`) WHERE cl.edition = %q RETURN count(*)", "mid-year-pop-age", "one-off")

	var count int64
	rowExtractor := func(r *bolt.Result) error {
		var ok bool
		count, ok = r.Data[0].(int64)
		if !ok {
			return errors.New("failed to cast result to int64")
		}
		return nil
	}

	err = db.QueryForResult(query, nil, rowExtractor)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if count != 1 {
		log.Fatal(errors.New("count in correct expected 1 but was " + strconv.FormatInt(count, 10)))
		os.Exit(1)
	}
	fmt.Println("success")
}
