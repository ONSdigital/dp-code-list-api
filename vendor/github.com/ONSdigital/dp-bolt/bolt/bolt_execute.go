package bolt

import "github.com/pkg/errors"

type Params map[string]interface{}


type Stmt struct {
	Query  string
	Params Params
}

func (d *DB) Exec(s Stmt) (int64, map[string]interface{}, error) {
	if s.Query == "" {
		return 0, nil, nil
	}

	conn, err := d.pool.OpenPool()
	if err != nil {
		return 0, nil, errors.WithMessage(err, "error opening neo4j connection")
	}
	defer conn.Close()

	res, err := conn.ExecNeo(s.Query, s.Params)
	if err != nil {
		return 0, nil, errors.WithMessage(err, "error executing statement")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, nil, errors.WithMessage(err, "error getting rows affected count from result")
	}
	return rowsAffected, res.Metadata(), nil
}
