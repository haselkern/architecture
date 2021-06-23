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

package usecase

import (
	flag "flag"
	fmt "fmt"
	os "os"
	strconv "strconv"
)

// MyConfig is use case feature flag configuration.
//
// The stereotype of this type is 'cfg'.
type MyConfig struct {
	// FancyFeature is the fancy feature toggle.
	FancyFeature bool
}

// Reset restores this instance to the default state.
//  * The default value of FancyFeature is 'false'
func (m *MyConfig) Reset() {
	m.FancyFeature = false
}

// ParseEnv tries to parse the environment variables into this instance.
// It will only set those values, which have been actually defined.
// If values cannot be parsed, an error is returned.
//  * FancyFeature is parsed from variable 'TICKETS_USECASE_FANCYFEATURE' if it has been set.
func (m *MyConfig) ParseEnv() error {
	if value, ok := os.LookupEnv("TICKETS_USECASE_FANCYFEATURE"); ok {
		parsed, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("unable to parse flag 'TICKETS_USECASE_FANCYFEATURE': %w", err)
		}

		m.FancyFeature = parsed
	}

	return nil
}

// ConfigureFlags configures the flags to be ready to get evaluated. The default values are taken from the struct at calling time.
// After calling, use flags.Parse() to load the values.
// The default values are the field values at calling time.
// Example:
//   cfg := MyConfig{}
//   cfg.Reset()
//   flags := flag.NewFlagSet(`my app`, flag.ExitOnError)
//   cfg.ConfigureFlags(flags)
//   flags.Parse(os.Args[1:])
//
// The following flags will be tied to this instance:
//   - there are no flags available.
func (m *MyConfig) ConfigureFlags(flags *flag.FlagSet) {
}
