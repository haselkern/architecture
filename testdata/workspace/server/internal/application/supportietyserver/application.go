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

package supportietyserver

import (
	context "context"
	errors "errors"
	flag "flag"
	fmt "fmt"
	core "github.com/golangee/architecture/testdata/workspace/server/internal/tickets/core"
	usecase "github.com/golangee/architecture/testdata/workspace/server/internal/tickets/usecase"
	os "os"
	filepath "path/filepath"
)

// Application embeds the defaultApplication to provide the default application behavior.
// It also provides the inversion of control injection mechanism for all bounded contexts.
type Application struct {
	defaultApplication
}

func NewApplication(ctx context.Context) (*Application, error) {
	a := &Application{}
	a.defaultApplication.self = a
	if err := a.init(ctx); err != nil {
		return nil, fmt.Errorf("cannot init application: %w", err)
	}

	return a, nil
}

// defaultApplication aggregates all contained bounded contexts and starts their driver adapters.
type defaultApplication struct {
	// cfg contains the global read-only configuration for all bounded contexts.
	cfg Configuration

	// self provides a pointer to the actual Application instance to provide
	// one level of a quasi-vtable calling indirection for simple method 'overriding'.
	self *Application

	ticketsUsecaseTickets *usecase.Tickets
}

// configure resets, prepares and parses the configuration. The priority of evaluation
func (d *defaultApplication) configure() error {
	usrCfgHome, err := os.UserConfigDir()
	if err == nil {
		usrCfgHome = filepath.Join(usrCfgHome, ".supportiety-server", "settings.json")
	}

	filename := flag.String("config", usrCfgHome, "filename to a configuration file in JSON format.")

	// prio 0: hardcoded defaults
	d.cfg.Reset()
	d.cfg.ConfigureFlags()
	flag.Parse()

	// prio 1: values from configuration file
	if *filename != "" {
		if err := d.cfg.ParseFile(*filename); err != nil {
			if *filename != usrCfgHome || !errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("cannot explicitly parse configuration file: %w", err)
			}
		}
	}

	// prio 2: values from environment variables
	if err := d.cfg.ParseEnv(); err != nil {
		return fmt.Errorf("cannot parse environment variables: %w", err)
	}

	// prio 3: values from TODO wrong order
	return nil
}

func (d *defaultApplication) init(ctx context.Context) error {
	if err := d.configure(); err != nil {
		return fmt.Errorf("cannot configure: %w", err)
	}

	return nil
}

func (_ defaultApplication) Run(ctx context.Context) error {
	return nil
}

func (d *defaultApplication) getTicketsUsecaseMyConfig() (usecase.MyConfig, error) {
	panic("assemble super config, parse that once and then poke from that")
}

func (d *defaultApplication) getTicketsCoreTickets() (core.Tickets, error) {
	panic("find different implementations and make them configurable, e.g. mysql vs postgres")
}

func (d *defaultApplication) getTicketsUsecaseTickets() (*usecase.Tickets, error) {
	if d.ticketsUsecaseTickets != nil {
		return d.ticketsUsecaseTickets, nil
	}

	myCfg, err := d.self.getTicketsUsecaseMyConfig()
	if err != nil {
		return nil, fmt.Errorf("cannot get parameter 'myCfg': %w", err)
	}

	tickets, err := d.self.getTicketsCoreTickets()
	if err != nil {
		return nil, fmt.Errorf("cannot get parameter 'tickets': %w", err)
	}

	s, err := usecase.NewTickets(myCfg, tickets)
	if err != nil {
		return nil, fmt.Errorf("cannot create service 'Tickets': %w", err)
	}

	d.ticketsUsecaseTickets = s

	return s, nil
}
