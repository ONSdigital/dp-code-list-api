package driver

import (
	"context"
	"errors"

	"github.com/ONSdigital/dp-code-list-api/models"
)

var ErrNotFound = errors.New("not found")

type Driver interface {
	// Open returns a new connection to the database.
	// The name is a string in a driver-specific format.
	//
	// Open may return a cached connection (one previously
	// closed), but doing so is unnecessary; the sql package
	// maintains a pool of idle connections for efficient re-use.
	//
	// The returned connection is only used by one goroutine at a
	// time.
	Open() (Conn, error)
	codelist
}

type codelist interface {
	GetCodeLists(ctx context.Context, filterBy string) (*models.CodeListResults, error)
	GetCodeList(ctx context.Context, codeListID string) (*models.CodeList, error)
	GetEditions(ctx context.Context, codeListID string) (*models.Editions, error)
	GetEdition(ctx context.Context, codeListID, edition string) (*models.Edition, error)
	GetCodes(ctx context.Context, codeListID, edition string) (*models.CodeResults, error)
	GetCode(ctx context.Context, codeListID, edition string, code string) (*models.Code, error)
	GetCodeDatasets(ctx context.Context, codeListID, edition string, code string) (*models.Datasets, error)
}

var ErrBadConn = errors.New("driver: bad connection")

type Pinger interface {
	Ping(ctx context.Context) error
}

type Conn interface {
	// // Prepare returns a prepared statement, bound to this connection.
	// Prepare(query string) (Stmt, error)

	// Close invalidates and potentially stops any current
	// prepared statements and transactions, marking this
	// connection as no longer in use.
	//
	// Because the sql package maintains a free pool of
	// connections and only calls Close when there's a surplus of
	// idle connections, it shouldn't be necessary for drivers to
	// do their own connection caching.
	Close() error
}

type Connector interface {
	// Connect returns a connection to the database.
	// Connect may return a cached connection (one previously
	// closed), but doing so is unnecessary; the sql package
	// maintains a pool of idle connections for efficient re-use.
	//
	// The provided context.Context is for dialing purposes only
	// (see net.DialContext) and should not be stored or used for
	// other purposes.
	//
	// The returned connection is only used by one goroutine at a
	// time.
	Connect(context.Context) (Conn, error)

	// Driver returns the underlying Driver of the Connector,
	// mainly to maintain compatibility with the Driver method
	// on sql.DB.
	Driver() Driver
}
