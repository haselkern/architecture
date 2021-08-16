package model

import "fmt"

// Validator should be implemented by all model types to allow for easy validation.
type Validator interface {
	// Validate should return nil when everything is okay and a descriptive error
	// otherwise. To validate references etc. a reference to the project is given.
	Validate(project *Project) error
}

// Project describes information provided for the whole project.
// TODO positional information in types
type Project struct {
	// You *need* to specify a version of architecture to use, but the version may be a wildcard,
	// which automatically will use the latest version of architecture.
	// This is according to https://semver.org/ spec.
	ArcVersion string `tadl:"arc_version"`
	// The name for this domain.
	Name string `tadl:"name"`
	// A short description for this domain.
	Description     string `tadl:"description"`
	BoundedContexts []BoundedContext
	Stories         []Story `tadl:"story"`
	Glossary        Glossary
}

// Validate will check the whole project for validity.
// Should everything be correct, then nil is returned.
// Should something be wrong an error describing the problem is returned.
func (p *Project) Validate(*Project) error {
	// Check if every method has a user story referenced.
	for _, boundedContext := range p.BoundedContexts {
		if err := boundedContext.Validate(p); err != nil {
			return fmt.Errorf("project '%s' invalid: %w", p.Name, err)
		}
	}

	return nil
}

type Glossary struct {
	Definitions map[string]string `tadl:",inner"`
}

type License struct {
	Name string `tadl:",inner"`
}

// BoundedContext describes meta information for a bounded context.
type BoundedContext struct {
	Name        string `tadl:"name"`
	Description string `tadl:"description"`
	// License is one of the identifiers from here: https://spdx.org/licenses/
	// This will allow the generator to download the corresponding license from here:
	// https://github.com/spdx/license-list-data
	License    License `tadl:"license"`
	Authors    Authors
	Generators GeneratorSelection `tadl:"generator"`
	Artifacts  Artifacts
}

func (b *BoundedContext) Validate(project *Project) error {
	if err := b.Artifacts.Validate(project); err != nil {
		return fmt.Errorf("bounded context '%s' invalid: %w", b.Name, err)
	}

	return nil
}

// Authors is list of authors that contributed to this context.
type Authors struct {
	Authors []Author `tadl:"author"`
}

// Author is a person with a name and a mail address.
type Author struct {
	Name string `tadl:"name"`
	Mail string `tadl:"mail"`
}

// GeneratorSelection can contain an arbitrary selection of types
// for which projects should be generated for a given BoundedContext.
type GeneratorSelection struct {
	Go      *GoGenerator      `tadl:"go"`
	Android *AndroidGenerator `tadl:"android"`
}

// GoGenerator is a generator to create go projects with.
type GoGenerator struct {
	// Package is the name of the go package that will be generated.
	Package string       `tadl:"package"`
	Build   DesktopBuild `tadl:"build"`
}

// DesktopBuild contains several build targets for desktop operating systems.
// Field Darwin could contain 'amd64' if we should build for a 64-bit Apple device.
type DesktopBuild struct {
	Darwin []string `tadl:"darwin"`
	Linux  []string `tadl:"linux"`
}

// AndroidGenerator does nothing and is only used as a demonstration of different
// generator backends.
type AndroidGenerator struct {
	Package string `tadl:"package"`
}