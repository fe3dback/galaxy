package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func (f *factory) provideRootDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(fmt.Sprintf("can`t provide root dir: %v", err))
	}

	return dir
}

func (f *factory) provideResourcesDirectory() string {
	return fmt.Sprintf("%s/resources", f.provideRootDirectory())
}

func (f *factory) provideFontsDirectory() string {
	return fmt.Sprintf("%s/fonts", f.provideResourcesDirectory())
}
