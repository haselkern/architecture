// Code generated by golangee/architecture. DO NOT EDIT.

package usecase

import (
	context "context"
	uuid "github.com/golangee/uuid"
)

// BookLoaningMock is a mock implementation of BookLoaning.
type BookLoaningMock struct {
	// RentFunc mocks the Rent function.
	RentFunc func(ctx context.Context, bookId uuid.UUID, userId uuid.UUID) error
	// CheckCustomerIdFunc mocks the CheckCustomerId function.
	CheckCustomerIdFunc func(ctx context.Context, userId uuid.UUID) error
}

// Rent loans a book.
//
// The parameter 'ctx' is the context to control timeouts and cancellations.
//
// The parameter 'bookId' is the id of the book.
//
// The parameter 'userId' is the id of the user, who loans the book.
//
// The result 'error' indicates a violation of pre- or invariants and represents an implementation specific failure.
func (m BookLoaningMock) Rent(ctx context.Context, bookId uuid.UUID, userId uuid.UUID) error {
	if m.RentFunc != nil {
		return m.RentFunc(ctx, bookId, userId)
	}

	panic("mock not available: Rent")
}

// CheckCustomerId validates if the user is registered and active.
//
// The parameter 'ctx' is the context to control timeouts and cancellations.
//
// The parameter 'userId' is the users id.
//
// The result 'error' indicates a violation of pre- or invariants and represents an implementation specific failure.
func (m BookLoaningMock) CheckCustomerId(ctx context.Context, userId uuid.UUID) error {
	if m.CheckCustomerIdFunc != nil {
		return m.CheckCustomerIdFunc(ctx, userId)
	}

	panic("mock not available: CheckCustomerId")
}
