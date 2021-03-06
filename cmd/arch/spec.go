package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

type (
	ComponentID  = string
	VendorID     = string
	RelativePath = string

	Spec struct {
		Path       string                        `yaml:"path"`
		Vendors    map[VendorID]SpecVendor       `yaml:"vendors"`
		Components map[ComponentID]SpecComponent `yaml:"components"`
		Deps       map[ComponentID]SpecContract  `yaml:"deps"`
		Common     []ComponentID                 `yaml:"common"`
		Exclude    []RelativePath                `yaml:"exclude"`
	}

	SpecComponent struct {
		In    string `yaml:"in"`
		Depth bool   `yaml:"depth"`
	}

	SpecVendor struct {
		In string `yaml:"in"`
	}

	SpecContract struct {
		MayDependOn []ComponentID `yaml:"mayDependOn"`
		CanUse      []VendorID    `yaml:"canUse"`
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
		Depth          bool
		AllowedDeps    []string
		AllowedVendors []string
	}
)

func parseSpec(path string) ParsedSpec {
	spec := Spec{}

	buffer, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		panic(fmt.Sprintf("can`t open spec %s: %v", path, err))
	}

	err = yaml.Unmarshal(buffer, &spec)
	if err != nil {
		panic(fmt.Sprintf("can`t parse spec: %v", err))
	}

	// assemble spec
	dirPath := filepath.Dir(path)
	parsed := ParsedSpec{
		Path:    spec.Path,
		Modules: map[string]Module{},
		Exclude: []string{
			cleanPath(fmt.Sprintf("%s/%s", dirPath, ".git")),
		},
	}

	for _, excludeRelativeDir := range spec.Exclude {
		parsed.Exclude = append(parsed.Exclude, cleanPath(fmt.Sprintf("%s/%s", dirPath, excludeRelativeDir)))
	}

	for id, component := range spec.Components {
		allowedDeps := make([]string, 0)
		allowedVendors := make([]string, 0)

		allowedDeps = append(allowedDeps, spec.Deps[id].MayDependOn...)
		allowedDeps = append(allowedDeps, spec.Common...)
		allowedVendors = append(allowedVendors, spec.Deps[id].CanUse...)

		deps := make([]string, 0)
		for _, depName := range allowedDeps {
			dep, ok := spec.Components[depName]
			if !ok {
				panic(fmt.Sprintf("can`t find component: %s", depName))
			}

			depAbsolutePath := cleanPath(fmt.Sprintf("%s/%s", dirPath, dep.In))
			deps = append(deps, fmt.Sprintf("%s/%s", spec.Path, dep.In))

			if dep.Depth {
				err = filepath.Walk(depAbsolutePath, func(path string, info os.FileInfo, err error) error {
					if !info.IsDir() {
						return nil
					}

					relativePath := strings.TrimPrefix(path, dirPath)
					deps = append(deps, fmt.Sprintf("%s%s", spec.Path, relativePath))

					return nil
				})
				if err != nil {
					panic(fmt.Sprintf("can`t walk for dep `%s`: %v", depAbsolutePath, err))
				}
			}
		}

		vendors := make([]string, 0)
		for _, vendorName := range allowedVendors {
			vendor, ok := spec.Vendors[vendorName]
			if !ok {
				panic(fmt.Sprintf("can`t find vendor: %s", vendorName))
			}

			vendors = append(vendors, vendor.In)
		}

		parsed.Modules[id] = Module{
			ID:             id,
			PathRelative:   component.In,
			PathAbsolute:   cleanPath(fmt.Sprintf("%s/%s", dirPath, component.In)),
			Depth:          component.Depth,
			AllowedDeps:    uniqueStrings(deps),
			AllowedVendors: uniqueStrings(vendors),
		}
	}

	return parsed
}

func cleanPath(p string) string {
	return filepath.Clean(p)
}

func uniqueStrings(list []string) []string {
	unique := make([]string, 0)
	check := make(map[string]struct{})

	for _, s := range list {
		if _, ok := check[s]; ok {
			continue
		}

		unique = append(unique, s)
		check[s] = struct{}{}
	}

	return unique
}

func (s ParsedSpec) getModule(path string) (Module, bool) {
	dir := filepath.Dir(path)

	wideModules := make([]Module, 0)

	for _, mod := range s.Modules {
		if mod.PathAbsolute == dir {
			return mod, true
		}

		if !mod.Depth {
			continue
		}

		wideModules = append(wideModules, mod)
	}

	sort.SliceStable(wideModules, func(i, j int) bool {
		return len(wideModules[j].PathAbsolute) < len(wideModules[i].PathAbsolute)
	})

	for _, mod := range wideModules {
		if strings.HasPrefix(dir, mod.PathAbsolute) {
			return mod, true
		}
	}

	return Module{}, false
}
