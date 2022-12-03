package harvest

import (
    "errors"
    "fmt"
    "go/ast"
    "go/parser"
    "go/token"
    "path/filepath"
    "strings"

    "github.com/qhzhyt/harvest/pkg/annotation"
    "github.com/qhzhyt/harvest/pkg/harvest/types"
)

func ParseMain(mainFilePath string) (*types.Harvest, error) {
    mainFilePath, err := filepath.Abs(mainFilePath)

    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    fset := token.NewFileSet()

    f, err := parser.ParseFile(fset, mainFilePath, nil, parser.ParseComments|parser.AllErrors)

    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    for _, decl := range f.Decls {
        funcDecl, ok := decl.(*ast.FuncDecl)

        if ok && funcDecl.Name.Name == "main" {
            for _, a := range annotation.Parse(funcDecl.Doc.Text()) {
                fmt.Println(a.Name, a.Args)
                if strings.ToLower(string(a.Name)) == "harvest" {
                    result := &types.Harvest{}
                    err = annotation.Unmarshal(a, result)
                    if err != nil {
                        return nil, err
                    }
                    return result, nil
                }
            }

        }
    }
    return nil, errors.New("main harvest not found")
}
