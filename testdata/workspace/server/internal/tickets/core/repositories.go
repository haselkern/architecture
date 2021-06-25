// Code generated by golangee/eearc; DO NOT EDIT.
//
// Copyright 2021 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	uuid "github.com/golangee/uuid"
	fs "io/fs"
	sync "sync"
)

// Tickets provides CRUD access to Tickets.
type Tickets interface {
	// CreateTicket creates a Ticket.
	//
	// The parameter id is the unique ticket id.
	// The result error if anything goes wrong.
	CreateTicket(id uuid.UUID) (Ticket, error)
}

// TicketRepo autogenerated repo
type TicketRepo interface {
}

// InMemoryTicketRepo implements a hashmap based in-memory implementation for Ticket entities.
type InMemoryTicketRepo struct {
	store map[uuid.UUID]Ticket
	mutex sync.RWMutex
}

// InsertOne inserts the entity or fails if already exists.
func (r InMemoryTicketRepo) InsertOne(entity Ticket) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.store[entity.ID]; ok {
		return fs.ErrExist
	}

	r.store[entity.ID] = entity

	return nil
}

// UpdateOne updates the entity or fails if does not exist.
func (r InMemoryTicketRepo) UpdateOne(entity Ticket) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.store[entity.ID]; !ok {
		return fs.ErrNotExist
	}

	r.store[entity.ID] = entity

	return nil
}

// FindOne finds the entity or fails if does not exist.
func (r InMemoryTicketRepo) FindOne(id uuid.UUID) (Ticket, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	v, ok := r.store[id]
	if !ok {
		return v, fs.ErrNotExist
	}

	return v, nil
}

// DeleteOne deletes the entity with the given id. Is a no-op if no such entity exists.
func (r InMemoryTicketRepo) DeleteOne(id uuid.UUID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	delete(r.store, id)

	return nil
}
