package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"text/template"
	"time"

	"github.com/iancoleman/strcase"
)

var flagEventTypesFile = flag.String("path", "", "path to event_type.go file with types const")

type eventTemplate struct {
	Name        string
	NameLC      string
	GeneratedAt string
}

type initTemplate struct {
	EventTypes  []string
	GeneratedAt string
}

var cmdDirectory string

func main() {
	flag.Parse()

	_, callerDir, _, ok := runtime.Caller(0)
	if !ok {
		panic("No cmdDirectory information")
	}
	cmdDirectory = path.Dir(callerDir)
	fmt.Printf("Working in '%s'\n", cmdDirectory)

	srcPath := getSource()
	fmt.Printf("parsing file: '%s'\n", srcPath)

	eventTypes := getEventTypes(srcPath)

	for _, eventType := range eventTypes {
		writeCode(srcPath, eventType, "generated")
		writeCode(srcPath, eventType, "dto")
	}

	// rewrite init
	filePath := fmt.Sprintf("%s/dispatcher_init.go", path.Dir(srcPath))
	initCode := renderInitTemplate(eventTypes)
	err := ioutil.WriteFile(filePath, initCode, 0644)
	checkFail(err, "can`t write init go file")
}

func writeCode(srcPath string, eventName string, suffix string) {
	filePath := fmt.Sprintf("%s/event_%s_%s.go", path.Dir(srcPath), strcase.ToSnake(eventName), suffix)
	templateName := fmt.Sprintf("event_%s.gohtml", suffix)

	created := generateFile(eventName, filePath, templateName)

	if created {
		fmt.Printf("[OK] File '%s' created\n", filePath)
	} else {
		fmt.Printf("[SKIP] File '%s' already exist\n", filePath)
	}
}

func generateFile(eventName string, filePath string, templateName string) bool {
	if _, err := os.Stat(filePath); err == nil {
		// skip
		return false
	}

	code := renderEventTemplate(eventName, templateName)
	err := ioutil.WriteFile(filePath, code, 0644)
	checkFail(err, fmt.Sprintf("can`t write go file '%s'", filePath))

	return true
}

func renderEventTemplate(eventName string, templateName string) []byte {
	templateFile := fmt.Sprintf("%s/%s", cmdDirectory, templateName)

	tmpl := template.New(templateName)
	tmpl, err := tmpl.ParseFiles(templateFile)
	checkFail(err, fmt.Sprintf("can`t open template '%s'", templateFile))

	var dest bytes.Buffer
	err = tmpl.Execute(&dest, eventTemplate{
		Name:        eventName,
		NameLC:      strcase.ToLowerCamel(eventName),
		GeneratedAt: time.Now().String(),
	})
	checkFail(err, fmt.Sprintf("can`t render template '%s'", templateFile))

	return dest.Bytes()
}

func renderInitTemplate(eventTypes []string) []byte {
	templateFile := fmt.Sprintf("%s/event_init.gohtml", cmdDirectory)

	tmpl := template.New(path.Base(templateFile))
	tmpl, err := tmpl.ParseFiles(templateFile)
	checkFail(err, fmt.Sprintf("can`t open template '%s'", templateFile))

	var dest bytes.Buffer
	err = tmpl.Execute(&dest, initTemplate{
		EventTypes:  eventTypes,
		GeneratedAt: time.Now().String(),
	})
	checkFail(err, fmt.Sprintf("can`t render template '%s'", templateFile))

	return dest.Bytes()
}

func getEventTypes(srcPath string) []string {
	fset := token.NewFileSet()
	fileAst, err := parser.ParseFile(fset, srcPath, nil, parser.AllErrors)
	checkFail(err, "can`t parse source")

	eventTypes := make([]string, 0)

	for _, d := range fileAst.Decls {
		switch decl := d.(type) {
		case *ast.GenDecl:
			if decl.Tok.String() != "const" {
				// want only const types
				continue
			}

			for _, s := range decl.Specs {
				switch spec := s.(type) {
				case *ast.ValueSpec:
					for _, name := range spec.Names {
						if !strings.HasPrefix(name.Name, "eventType") {
							continue
						}

						eventType := strings.TrimPrefix(name.Name, "eventType")
						eventTypes = append(eventTypes, eventType)
					}
				}
			}
		}
	}
	return eventTypes
}

func getSource() string {
	realPath := strings.Trim(*flagEventTypesFile, "\"")
	srcPath := path.Clean(realPath)

	fileInfo, err := os.Stat(srcPath)
	if os.IsNotExist(err) {
		failf(fmt.Sprintf("source file not found in '%s'", realPath))
	}
	checkFail(err, fmt.Sprintf("failed to check source file '%s'", realPath))

	if fileInfo.IsDir() {
		failf(fmt.Sprintf("source path `%s` should by file (not dir)", realPath))
	}

	return srcPath
}

func failf(explain string) {
	panic(explain)
}

func checkFail(err error, explain string) {
	if err == nil {
		return
	}

	panic(fmt.Sprintf("%s: %v", explain, err))
}
