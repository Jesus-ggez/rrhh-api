package data

import (
    "context"
    "database/sql"
    "log"
    "time"

    // driver
    _ "github.com/tursodatabase/libsql-client-go/libsql"
)

func NewPool(token, url string) (*sql.DB, error) {
    dbPath := url + "?authToken=" + token

    p, err := sql.Open("libsql", dbPath)
    if err != nil {
        return nil, err
    }
    p.SetConnMaxLifetime(5 * time.Minute)
    p.SetMaxOpenConns(10)
    p.SetMaxIdleConns(5)

    ctx := context.Background()
    if err := p.PingContext(ctx); err != nil {
        return nil, err
    }
    log.Printf("Database pool connected successfully\n%s", p.Stats())
    return p, nil
}
/*
// i use only this
func Open(driverName, dataSourceName string) (*DB, error)
func (db *DB) SetConnMaxIdleTime(d time.Duration)
func (db *DB) SetConnMaxLifetime(d time.Duration)
func (db *DB) PingContext(ctx context.Context) error
func (db *DB) Ping() error // ...?
func (db *DB) SetMaxIdleConns(n int)
func (db *DB) SetMaxOpenConns(n int)
func (db *DB) Stats() DBStats

// this not in this instant
func OpenDB(c driver.Connector) *DB
func (db *DB) Begin() (*Tx, error)
func (db *DB) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)

func (db *DB) Close() error
func (db *DB) Conn(ctx context.Context) (*Conn, error)
func (db *DB) Driver() driver.Driver
func (db *DB) Exec(query string, args ...any) (Result, error)
func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (Result, error)


func (db *DB) Prepare(query string) (*Stmt, error)
func (db *DB) PrepareContext(ctx context.Context, query string) (*Stmt, error)
func (db *DB) Query(query string, args ...any) (*Rows, error)
func (db *DB) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)
func (db *DB) QueryRow(query string, args ...any) *Row
func (db *DB) QueryRowContext(ctx context.Context, query string, args ...any) *Row
*/

