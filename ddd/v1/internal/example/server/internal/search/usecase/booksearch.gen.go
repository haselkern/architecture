// Code generated by golangee/architecture. DO NOT EDIT.

package usecase

import (
	context "context"
	core "example-server/internal/search/core"
	uuid "github.com/golangee/uuid"
)

// BookSearch provides all user stories involved in searching books.
//
// The following user stories are covered:
//
//   * As a searcher, I want to search for keywords, so that I must not know the title or author.
//   * As a searcher, I want to have an autocomplete so that I get support while typing my keywords.
//   * As a searcher, I want to see book details, because I need to proof the relevance of the result.
//   * As a book admin, I want to change a title, because the book has a typo.
type BookSearch interface {
	// FindByTags searches for tags only.
	//
	// The parameter 'ctx' is the context to control timeouts and cancellations.
	//
	// The parameter 'query' provides tokens to search for, separated by spaces or commas.
	//
	// The result '[]Book' is a list of books which match to the query.
	//
	// The result 'error' indicates a violation of pre- or invariants and represents an implementation specific failure.
	FindByTags(ctx context.Context, query string) ([]core.Book, error)
	// Autocomplete proposes autocompletion values.
	//
	// The parameter 'ctx' is the context to control timeouts and cancellations.
	//
	// The parameter 'text' is the text to autocomplete.
	//
	// The result '[]AutoCompleteValue' is a list of proposals.
	//
	// The result 'error' indicates a violation of pre- or invariants and represents an implementation specific failure.
	Autocomplete(ctx context.Context, text string) ([]AutoCompleteValue, error)
	// Details returns the details of a book.
	//
	// The parameter 'ctx' is the context to control timeouts and cancellations.
	//
	// The parameter 'id' is the Id of a book.
	//
	// The parameter 'user' is the authenticated user.
	//
	// The result 'Book' is the according book.
	//
	// The result 'error' indicates a violation of pre- or invariants and represents an implementation specific failure.
	Details(ctx context.Context, id uuid.UUID, user AuthUser) (core.Book, error)
	// ChangeBookTitle changes the book title.
	//
	// The parameter 'titleModel' is to short.
	//
	// The result 'Book' the updated book.
	//
	// The result 'error' indicates a violation of pre- or invariants and represents an implementation specific failure.
	ChangeBookTitle(titleModel BookTitleSpec) (core.Book, error)
}

// AutoCompleteValue represents an auto completed value.
type AutoCompleteValue struct {
	// Value is the value to complete.
	Value string `json:"value"`
	// Score the probability of importance.
	Score float32 `json:"score"`
	// Synonyms alternative search suggestions.
	Synonyms []string `json:"synonyms"`
}

// AuthUser represents an authenticated user.
type AuthUser struct {
	// Age is the age of a user and determines if he can view the book or not.
	Age int `json:"age"`
}

// BookTitleSpec is for changing book titles.
type BookTitleSpec struct {
	// Title is a title.
	Title string `json:"title,omitempty"`
}
