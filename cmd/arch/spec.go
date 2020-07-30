package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type (
	ComponentID = string

	Spec struct {
		Path       string
		Components map[ComponentID]SpecComponent
		Common     []ComponentID
		Deps       map[ComponentID]string
		Exclude    []string
	}

	SpecComponent struct {
		In    string
		Inner bool
	}

	ParsedSpec struct {
		Path    string
		Modules map[string]Module
		Exclude []string
	}

	Module struct {
		ID             string
		PathRelative   string
		PathAbsolute   string
		IncludeSubDirs bool
		AllowedDeps    []string
	}
)

func parseSpec(path string) ParsedSpec {
	spec := Spec{}

	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("can`t open spec %s: %v", path, err))
	}

	err = yaml.Unmarshal(buffer, &spec)
	if err != nil {
		panic(fmt.Sprintf("can`t parse spec: %v", err))
	}

	// assemble spec

	parsed := ParsedSpec{
		Path:    spec.Path,
		Modules: map[string]Module{},
		Exclude: []string{},
	}
	dirPath := filepath.Dir(path)

	for _, excludeRelativeDir := range spec.Exclude {
		parsed.Exclude = append(parsed.Exclude, fmt.Sprintf("%s/%s", dirPath, excludeRelativeDir))
	}

	for id, component := range spec.Components {
		allowedDeps := make([]string, 0)

		if deps, ok := spec.Deps[id]; ok {
			deps = strings.TrimSpace(deps)
			deps = strings.TrimPrefix(deps, "->")

			for _, depName := range strings.Split(deps, ",") {
				allowedDeps = append(allowedDeps, strings.TrimSpace(depName))
			}
		}

		for _, depName := range spec.Common {
			allowedDeps = append(allowedDeps, strings.TrimSpace(depName))
		}

		deps := make([]string, 0)
		for _, depName := range allowedDeps {
			dep, ok := spec.Components[depName]
			if !ok {
				panic(fmt.Sprintf("can`t find component: %s", depName))
			}

			deps = append(deps, fmt.Sprintf("%s/%s", spec.Path, dep.In))
		}

		parsed.Modules[id] = Module{
			ID:             id,
			PathRelative:   component.In,
			PathAbsolute:   fmt.Sprintf("%s/%s", dirPath, component.In),
			IncludeSubDirs: component.Inner,
			AllowedDeps:    deps,
		}
	}

	return parsed
}

func (s ParsedSpec) getModule(path string) (Module, bool) {
	dir := filepath.Dir(path)

	for _, mod := range s.Modules {
		if mod.PathAbsolute == dir {
			return mod, true
		}

		if strings.HasPrefix(dir, mod.PathAbsolute) {
			return mod, true
		}
	}

	return Module{}, false
}
