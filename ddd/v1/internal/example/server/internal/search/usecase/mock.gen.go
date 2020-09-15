// Code generated by golangee/architecture. DO NOT EDIT.

package usecase

import (
	context "context"
	core "example-server/internal/search/core"
	uuid "github.com/golangee/uuid"
)

// BookSearchMock is a mock implementation of BookSearch.
type BookSearchMock struct {
	// FindByTagsFunc mocks the FindByTags function.
	FindByTagsFunc func(ctx context.Context, query string) ([]core.Book, error)
	// AutocompleteFunc mocks the Autocomplete function.
	AutocompleteFunc func(ctx context.Context, text string) ([]AutoCompleteValue, error)
	// DetailsFunc mocks the Details function.
	DetailsFunc func(ctx context.Context, id uuid.UUID) (core.Book, error)
	// ChangeBookTitleFunc mocks the ChangeBookTitle function.
	ChangeBookTitleFunc func(titleModel BookTitleSpec) (core.Book, error)
}

// FindByTags searches for tags only.
//
// The parameter 'ctx' is the context to control timeouts and cancellations.
//
// The parameter 'query' provides tokens to search for, separated by spaces or commas.
//
// The result '[]Book' is a list of books which match to the query.
//
// The result 'error' indicates a violation of pre- or invariants and represents an implementation specific failure.
func (m BookSearchMock) FindByTags(ctx context.Context, query string) ([]core.Book, error) {
	if m.FindByTagsFunc != nil {
		return m.FindByTagsFunc(ctx, query)
	}

	panic("mock not available: FindByTags")
}

// Autocomplete proposes autocompletion values.
//
// The parameter 'ctx' is the context to control timeouts and cancellations.
//
// The parameter 'text' is the text to autocomplete.
//
// The result '[]AutoCompleteValue' is a list of proposals.
//
// The result 'error' indicates a violation of pre- or invariants and represents an implementation specific failure.
func (m BookSearchMock) Autocomplete(ctx context.Context, text string) ([]AutoCompleteValue, error) {
	if m.AutocompleteFunc != nil {
		return m.AutocompleteFunc(ctx, text)
	}

	panic("mock not available: Autocomplete")
}

// Details returns the details of a book.
//
// The parameter 'ctx' is the context to control timeouts and cancellations.
//
// The parameter 'id' is the Id of a book.
//
// The result 'Book' is the according book.
//
// The result 'error' indicates a violation of pre- or invariants and represents an implementation specific failure.
func (m BookSearchMock) Details(ctx context.Context, id uuid.UUID) (core.Book, error) {
	if m.DetailsFunc != nil {
		return m.DetailsFunc(ctx, id)
	}

	panic("mock not available: Details")
}

// ChangeBookTitle changes the book title.
//
// The parameter 'titleModel' is to short.
//
// The result 'Book' the updated book.
//
// The result 'error' indicates a violation of pre- or invariants and represents an implementation specific failure.
func (m BookSearchMock) ChangeBookTitle(titleModel BookTitleSpec) (core.Book, error) {
	if m.ChangeBookTitleFunc != nil {
		return m.ChangeBookTitleFunc(titleModel)
	}

	panic("mock not available: ChangeBookTitle")
}
