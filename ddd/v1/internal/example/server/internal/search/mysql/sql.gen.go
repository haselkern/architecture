// Code generated by golangee/architecture. DO NOT EDIT.

// Package mysql contains specific repository implementations (aka SPI or driven adapter) for the mysql dialect.
// The repository is defined at the core layer (aka domain API).
package mysql

import (
	context "context"
	core "example-server/internal/search/core"
)

// MysqlBookRepository is an implementation of the core.BookRepository defined as SPI/driven port in the domain/core layer.
// The queries are specific for the mysql dialect.
type MysqlBookRepository struct {
}

// ReadAll returns all books.
//
// The parameter 'ctx' is the context to control timeouts and cancellations.
//
// The parameter 'offset' is the offset to return the entries for paging.
//
// The parameter 'limit' is the maximum amount of entries to return.
//
// The result '[]Book' is the list of books.
//
// The result 'error' indicates a violation of pre- or invariants and represents an implementation specific failure.
func (b *MysqlBookRepository) ReadAll(ctx context.Context, offset int64, limit int64) ([]core.Book, error) {
	const s = "SELECT * FROM book LIMIT ? OFFSET ?"
}

// Count enumerates all stored elements.
//
// The parameter 'ctx' is the context to control timeouts and cancellations.
//
// The result 'int64' is the actual count.
//
// The result 'error' indicates a violation of pre- or invariants and represents an implementation specific failure.
func (b *MysqlBookRepository) Count(ctx context.Context) (int64, error) {
	const s = "SELECT count(*) FROM book"
}

// FindOne finds exactly one entry.
//
// The parameter 'ctx' is the context to control timeouts and cancellations.
//
// The parameter 'dto' is the data transfer object to read into.
//
// The result 'error' indicates a violation of pre- or invariants and represents an implementation specific failure.
func (b *MysqlBookRepository) FindOne(ctx context.Context, dto *core.Book) error {
	const s = "SELECT * FROM book WHERE ?"
}
