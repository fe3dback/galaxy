package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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
	err := filepath.Walk(dirRoot, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}

		for _, excludeDir := range spec.Exclude {
			if path == excludeDir {
				return nil
			}

			if strings.HasPrefix(path, excludeDir) {
				return nil
			}
		}

		packages, err := parser.ParseDir(fSet, path, fileFilter, parser.ImportsOnly)
		if err != nil {
			panic(fmt.Sprintf("can`t parse dir `%s`: %v", path, err))
		}

		warnings = append(warnings, checkPackages(spec, packages)...)
		return nil
	})
	if err != nil {
		panic(fmt.Sprintf("can`t walk: %v", err))
	}

	fmt.Println("--------------------------------")

	if len(warnings) == 0 {
		fmt.Println("[OK] No deps warning found!")
		return
	}

	for _, warning := range warnings {
		fmt.Printf("[WARN] %s\n\n", warning)
	}
}

func checkPackages(spec ParsedSpec, packages map[string]*ast.Package) []string {
	warnings := make([]string, 0)

	for _, pack := range packages {
		fmt.Println(pack.Name)

		for goFileName, goFile := range pack.Files {
			fmt.Printf("- %s\n", goFileName)
			module, ok := spec.getModule(goFileName)
			if !ok {
				fmt.Printf("[!] Can`t find module for file: %s\n", goFileName)
				continue
			}

			for _, goImport := range goFile.Imports {
				importPath := strings.Trim(goImport.Path.Value, "\"")

				if !strings.HasPrefix(importPath, spec.Path) {
					continue
				}

				if !inList(importPath, module.AllowedDeps) {
					warnings = append(warnings, fmt.Sprintf(
						"Module `%s`, file `%s` should not depend on `%s`, may depend only on %v",
						module.ID,
						goFileName,
						importPath,
						module.AllowedDeps,
					))
				}

				fmt.Printf("    [imp] %s\n", importPath)
			}

			for _, dep := range module.AllowedDeps {
				fmt.Printf("    [dep] %s\n", dep)
			}

		}
	}

	return warnings
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
