package annotation

import (
    "fmt"
    "go/ast"
    "go/parser"
    "go/token"
    "os"
    "path"
    "path/filepath"
)

func ScanLocalPackage(modulePath string, pkg string, deepScan bool) (map[string][]*Annotation, error) {

    result := map[string][]*Annotation{}

    var curAnnotations []*Annotation

    pkgPath, err := filepath.Abs(path.Join(modulePath, pkg))

    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    fset := token.NewFileSet()

    fs, err := parser.ParseDir(fset, pkgPath, nil, parser.ParseComments|parser.AllErrors)

    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    //ast.Print(fset, fs)

    for pn, f := range fs {

        for fn, file := range f.Files {
            for _, decl := range file.Decls {

                switch d := decl.(type) {
                case *ast.FuncDecl:
                    target := &Target{
                        Module:      modulePath,
                        Package:     pkg,
                        Type:        TargetTypeFunc,
                        Name:        d.Name.Name,
                        FuncRecv:    "",
                        Annotations: nil,
                    }
                    if d.Recv != nil {

                        switch rt := d.Recv.List[0].Type.(type) {
                        case *ast.StarExpr:
                            target.FuncRecv = "*" + rt.X.(*ast.Ident).Name
                        case *ast.Ident:
                            target.FuncRecv = rt.Name
                        }
                        //fmt.Println(a.FuncRecv, a.Name)
                    }
                    for _, a := range Parse(d.Doc.Text()) {
                        a.File = fn
                        a.Package = pn
                        a.Target.Type = TargetTypeFunc
                        a.Target.Name = d.Name.Name
                        a.Target.Module = modulePath
                        a.Target.Package = pkg

                        curAnnotations = append(curAnnotations, a)
                    }
                case *ast.GenDecl:
                    annotations := Parse(d.Doc.Text())
                    for _, spec := range d.Specs {
                        for _, annotation := range annotations {

                            annotation.File = fn
                            annotation.Package = pn
                            annotation.Target.Module = modulePath
                            annotation.Target.Package = pkg

                            switch s := spec.(type) {

                            case *ast.TypeSpec:

                                annotation.Target.Name = s.Name.Name

                                switch s.Type.(type) {
                                case *ast.StructType:
                                    annotation.Target.Type = TargetTypeStruct
                                case *ast.InterfaceType:
                                    annotation.Target.Type = TargetTypeInterface
                                //case *ast.FuncType:
                                //case *ast.:
                                default:

                                    fmt.Println("TypeSpec: ", s.Type)
                                    continue

                                }

                            case *ast.ValueSpec:
                                annotation.Target.Name = s.Names[0].Name
                                annotation.Target.Type = TargetTypeVar
                                annotation.Target.Module = modulePath
                                annotation.Target.Package = pkg
                            default:
                                fmt.Println("Spec: ", s)
                            }

                            curAnnotations = append(curAnnotations, annotation)

                        }
                    }
                default:
                    fmt.Println("Decl: ", d)
                }
            }

        }

    }

    if deepScan {
        entries, err := os.ReadDir(pkgPath)
        if err != nil {
            return nil, err
        }

        for _, entry := range entries {
            if entry.IsDir() {
                annotations, err := ScanLocalPackage(modulePath, path.Join(pkg, entry.Name()), true)
                if err != nil {
                    return nil, err
                }

                for pkg0, ans := range annotations {
                    result[pkg0] = ans
                }
            }
        }
    }

    result[pkgPath] = curAnnotations

    return result, nil
}
