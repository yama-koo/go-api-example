package infrastructure

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yama-koo/go-api-example/interface/database"
)

// SQLHandler struct
type SQLHandler struct {
	Conn *sql.DB
}

// NewSQLHandler function
func NewSQLHandler() *SQLHandler {
	conn, err := sql.Open("mysql", "root:@tcp(db:3306/sample")
	if err != nil {
		panic(err.Error)
	}

	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}

// Execute function
func (handler *SQLHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SQLResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}

	res.Result = result
	return res, nil
}

// Query function
func (handler *SQLHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SQLRow), err
	}

	row := new(SQLRow)
	row.Rows = rows
	return row, nil
}

// SQLResult struct
type SQLResult struct {
	Result sql.Result
}

// LastInsertId function
func (r SQLResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

// RowsAffected function
func (r SQLResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

// SQLRow struct
type SQLRow struct {
	Rows *sql.Rows
}

// Scan function
func (r SQLRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest)
}

// Next function
func (r SQLRow) Next() bool {
	return r.Rows.Next()
}

// Close function
func (r SQLRow) Close() error {
	return r.Rows.Close()
}
