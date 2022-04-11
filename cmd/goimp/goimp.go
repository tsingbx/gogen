package main

import (
	"fmt"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"runtime"

	"golang.org/x/tools/go/gcexportdata"
)

func main() {
	val := filepath.Join(runtime.GOROOT(), "pkg/mod")
	os.Setenv("GOMODCACHE", val)
	fmt.Println("GOMODCACHE:", val)

	fset := token.NewFileSet()
	packages := make(map[string]*types.Package)
	imp := gcexportdata.NewImporter(fset, packages)

	_, err := imp.Import("go/types")
	fmt.Println("Import result:", err)

	_, err = imp.Import("golang.org/x/tools/go/gcexportdata")
	fmt.Println("Import result:", err)
}
