package main

//go:generate go run $GOFILE

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
	"time"

	"github.com/iancoleman/strcase"
)

type (
	Data struct {
		GeneratedAt    string
		ResourceGroups []ResourceGroup
		TemplateFile   string
		DestFile       string
	}
	ResourceGroup struct {
		Name      string
		Resources []Resource
	}
	Resource struct {
		Name string
		Path string
	}
)

func main() {
	data := Data{
		GeneratedAt: time.Now().UTC().String(),
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	dirCurrent := filepath.Dir(filename)
	dirRoot := filepath.Dir(filepath.Dir(dirCurrent))
	dirGen := filepath.Join(dirRoot, "generated")
	dirResources := filepath.Join(dirRoot, "resources")

	data.TemplateFile = filepath.Join(dirCurrent, "template.gohtml")
	data.DestFile = filepath.Join(dirGen, "resources_path.go")

	dirs, err := ioutil.ReadDir(dirResources)
	if err != nil {
		panic(err)
	}

	groups := make([]ResourceGroup, 0)

	for _, rs := range dirs {
		if !rs.IsDir() {
			panic(fmt.Sprintf("not expected resource file `%s` in root resources dir", rs.Name()))
		}

		// iterate files
		resources := make([]Resource, 0)

		filesDir := filepath.Join(dirResources, rs.Name())

		err = filepath.Walk(filesDir, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			relativePath := strings.TrimPrefix(path, dirRoot+"/")
			resName := strings.TrimSuffix(relativePath, filepath.Ext(info.Name()))
			resName = strings.ReplaceAll(resName, "/", "_")
			resName = strcase.ToCamel(resName)

			resources = append(resources, Resource{
				Name: resName,
				Path: relativePath,
			})

			return nil
		})
		if err != nil {
			panic(err)
		}

		groups = append(groups, ResourceGroup{
			Name:      strings.Title(rs.Name()),
			Resources: resources,
		})
	}

	data.ResourceGroups = groups
	render(data)
}

func render(data Data) {
	tmpl := template.New(filepath.Base(data.TemplateFile))

	tmpl, err := tmpl.ParseFiles(data.TemplateFile)
	if err != nil {
		panic(err)
	}

	dest, err := os.Create(data.DestFile)
	if err != nil {
		panic(err)
	}
	defer dest.Close()

	err = tmpl.Execute(dest, data)
	if err != nil {
		panic(err)
	}
}
