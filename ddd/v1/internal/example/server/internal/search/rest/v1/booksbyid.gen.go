// Code generated by golangee/architecture. DO NOT EDIT.

package rest

import (
	uuid "github.com/golangee/uuid"
	http "net/http"
)

// BooksByIdGetContext provides the specific http request and response context including already parsed parameters.
type BooksByIdGetContext struct {
	// Request contains the raw http request.
	Request *http.Request
	// Writer contains a reference to the raw http response writer.
	Writer http.ResponseWriter
	// ClientId contains the parsed header parameter for 'clientId'.
	ClientId string
}

// BooksByIdDeleteContext provides the specific http request and response context including already parsed parameters.
type BooksByIdDeleteContext struct {
	// Request contains the raw http request.
	Request *http.Request
	// Writer contains a reference to the raw http response writer.
	Writer http.ResponseWriter
	// ClientId contains the parsed header parameter for 'clientId'.
	ClientId string
}

// BooksByIdPutContext provides the specific http request and response context including already parsed parameters.
type BooksByIdPutContext struct {
	// Request contains the raw http request.
	Request *http.Request
	// Writer contains a reference to the raw http response writer.
	Writer http.ResponseWriter
	// ClientId contains the parsed header parameter for 'clientId'.
	ClientId string
}

// BooksByIdPostContext provides the specific http request and response context including already parsed parameters.
type BooksByIdPostContext struct {
	// Request contains the raw http request.
	Request *http.Request
	// Writer contains a reference to the raw http response writer.
	Writer http.ResponseWriter
	// ClientId contains the parsed header parameter for 'clientId'.
	ClientId string
	// Id contains the parsed path parameter for 'id'.
	Id uuid.UUID
	// Bearer contains the parsed header parameter for 'bearer'.
	Bearer string
	// XSpecialSomething contains the parsed header parameter for 'x-special-something'.
	XSpecialSomething string
	// Offset contains the parsed query parameter for 'offset'.
	Offset int64
	// Limit contains the parsed query parameter for 'limit'.
	Limit int64
}

// BooksById represents the REST resource api/v1/books/:id.
// Resource to manage a single book.
type BooksById interface {
	// GetBooksById represents the http GET request on the /books/:id resource.
	// Returns a single book.
	GetBooksById(ctx BooksByIdGetContext) error
	// DeleteBooksById represents the http DELETE request on the /books/:id resource.
	// Removes a single book.
	DeleteBooksById(ctx BooksByIdDeleteContext) error
	// PutBooksById represents the http PUT request on the /books/:id resource.
	// Updates a book.
	PutBooksById(ctx BooksByIdPutContext) error
	// PostBooksById represents the http POST request on the /books/:id resource.
	// Creates a new book.
	PostBooksById(ctx BooksByIdPostContext) error
}

// BooksByIdMock is a mock implementation of BooksById.
// BooksById represents the REST resource api/v1/books/:id.
// Resource to manage a single book.
type BooksByIdMock struct {
	// GetBooksByIdFunc mocks the GetBooksById function.
	GetBooksByIdFunc func(ctx BooksByIdGetContext) error
	// DeleteBooksByIdFunc mocks the DeleteBooksById function.
	DeleteBooksByIdFunc func(ctx BooksByIdDeleteContext) error
	// PutBooksByIdFunc mocks the PutBooksById function.
	PutBooksByIdFunc func(ctx BooksByIdPutContext) error
	// PostBooksByIdFunc mocks the PostBooksById function.
	PostBooksByIdFunc func(ctx BooksByIdPostContext) error
}

// GetBooksById represents the http GET request on the /books/:id resource.
// Returns a single book.
func (m BooksByIdMock) GetBooksById(ctx BooksByIdGetContext) error {
	if m.GetBooksByIdFunc != nil {
		return m.GetBooksByIdFunc(ctx)
	}

	panic("mock not available: GetBooksById")
}

// DeleteBooksById represents the http DELETE request on the /books/:id resource.
// Removes a single book.
func (m BooksByIdMock) DeleteBooksById(ctx BooksByIdDeleteContext) error {
	if m.DeleteBooksByIdFunc != nil {
		return m.DeleteBooksByIdFunc(ctx)
	}

	panic("mock not available: DeleteBooksById")
}

// PutBooksById represents the http PUT request on the /books/:id resource.
// Updates a book.
func (m BooksByIdMock) PutBooksById(ctx BooksByIdPutContext) error {
	if m.PutBooksByIdFunc != nil {
		return m.PutBooksByIdFunc(ctx)
	}

	panic("mock not available: PutBooksById")
}

// PostBooksById represents the http POST request on the /books/:id resource.
// Creates a new book.
func (m BooksByIdMock) PostBooksById(ctx BooksByIdPostContext) error {
	if m.PostBooksByIdFunc != nil {
		return m.PostBooksByIdFunc(ctx)
	}

	panic("mock not available: PostBooksById")
}
