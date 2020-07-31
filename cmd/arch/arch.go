package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

type ImportType uint8

const (
	ImportTypeStdLib ImportType = iota
	ImportTypeProject
	ImportTypeVendor
)

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	dirCurrent := filepath.Dir(filename)
	dirRoot := filepath.Dir(filepath.Dir(dirCurrent))
	fmt.Printf("parse go in %s\n", dirRoot)

	spec := parseSpec(fmt.Sprintf("%s/arch.yml", dirRoot))
	warnings := make([]string, 0)

	fSet := token.NewFileSet()
	checkDirs := getProjectDirectories(dirRoot, spec.Exclude)

	for _, path := range checkDirs {
		packages, err := parser.ParseDir(fSet, path, fileFilter, parser.ImportsOnly)
		if err != nil {
			panic(fmt.Sprintf("can`t parse dir `%s`: %v", path, err))
		}

		warnings = append(warnings, checkPackages(dirRoot, spec, packages)...)
	}

	fmt.Println("--------------------------------")

	if len(warnings) == 0 {
		fmt.Println("[OK] No deps warning found!")
		return
	}

	for _, warning := range warnings {
		fmt.Printf("[WARN] %s\n\n", warning)
	}

	fmt.Printf("Total: %d warnings\n", len(warnings))
	os.Exit(1)
}

func checkPackages(dirRoot string, spec ParsedSpec, packages map[string]*ast.Package) []string {
	warnings := make([]string, 0)

	for _, pack := range packages {
		for goFileName, goFile := range pack.Files {
			module, ok := spec.getModule(goFileName)
			if !ok {
				warnings = append(warnings, fmt.Sprintf("Can`t find module for file: %s", goFileName))
				continue
			}

			for _, goImport := range goFile.Imports {
				importPath := strings.Trim(goImport.Path.Value, "\"")

				switch getImportType(importPath, spec) {
				case ImportTypeVendor:
					if !inList(importPath, module.AllowedVendors) {
						warnings = append(warnings, fmt.Sprintf(
							"Module `%s`, file `%s` can`t depend on vendor `%s`",
							module.ID,
							strings.TrimPrefix(goFileName, dirRoot),
							importPath,
						))
						continue
					}
				case ImportTypeProject:
					if !inList(importPath, module.AllowedDeps) {
						depsWarnings := ""
						for _, dep := range module.AllowedDeps {
							depsWarnings += fmt.Sprintf("\n - %s", strings.TrimPrefix(dep, spec.Path))
						}

						warnings = append(warnings, fmt.Sprintf(
							"Module `%s`, file `%s` should not depend on `%s`, may depend only on: %s",
							module.ID,
							strings.TrimPrefix(goFileName, dirRoot),
							strings.TrimPrefix(importPath, spec.Path),
							depsWarnings,
						))
					}
				}
			}
		}
	}

	return warnings
}

func getImportType(importPath string, spec ParsedSpec) ImportType {
	if !strings.Contains(importPath, ".") {
		return ImportTypeStdLib
	}

	if strings.HasPrefix(importPath, spec.Path) {
		return ImportTypeProject
	}

	return ImportTypeVendor
}

func getProjectDirectories(dirRoot string, exclude []string) []string {
	checkDirs := make([]string, 0)

	err := filepath.Walk(dirRoot, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}

		for _, excludeDir := range exclude {
			if path == excludeDir {
				return nil
			}

			if strings.HasPrefix(path, excludeDir) {
				return nil
			}
		}

		checkDirs = append(checkDirs, path)

		return nil
	})

	if err != nil {
		panic(fmt.Sprintf("can`t walk: %v", err))
	}

	sort.SliceStable(checkDirs, func(i, j int) bool {
		return len(checkDirs[j]) < len(checkDirs[i])
	})

	return checkDirs
}

func fileFilter(file os.FileInfo) bool {
	return !file.IsDir() && filepath.Ext(file.Name()) == ".go"
}

func inList(needle string, haystack []string) bool {
	for _, s := range haystack {
		if needle == s {
			return true
		}
	}

	return false
}
