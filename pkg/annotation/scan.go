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

func ScanLocalPackage(localPath string, modulePath string, pkg string, deepScan bool) (map[string][]*Target, error) {

	result := map[string][]*Target{}

	//var curAnnotations []*Annotation

	var targets []*Target

	pkgPath, err := filepath.Abs(path.Join(localPath, pkg))

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

	for _, f := range fs {

		for fn, file := range f.Files {
			for _, decl := range file.Decls {

				switch d := decl.(type) {
				case *ast.FuncDecl:
					target := &Target{
						File:        fn,
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
					target.Annotations = Parse(d.Doc.Text())

					targets = append(targets, target)

				case *ast.GenDecl:
					annotations := Parse(d.Doc.Text())

					for _, spec := range d.Specs {

						target := &Target{
							File:        fn,
							Module:      modulePath,
							Package:     pkg,
							Annotations: annotations,
						}

						switch s := spec.(type) {

						case *ast.TypeSpec:

							target.Name = s.Name.Name

							switch s.Type.(type) {
							case *ast.StructType:
								target.Type = TargetTypeStruct
							case *ast.InterfaceType:
								target.Type = TargetTypeInterface
							//case *ast.FuncType:
							//case *ast.:
							default:
								fmt.Println("TypeSpec: ", s.Type)
								continue
							}

						case *ast.ValueSpec:
							target.Name = s.Names[0].Name
							target.Type = TargetTypeVar
						default:
							fmt.Println("Spec: ", s)
						}

						targets = append(targets, target)
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
				annotations, err := ScanLocalPackage(localPath, modulePath, path.Join(pkg, entry.Name()), true)
				if err != nil {
					return nil, err
				}

				for pkg0, ans := range annotations {
					result[pkg0] = ans
				}
			}
		}
	}

	for _, target := range targets {
		target.ID = GenTargetID(target)
		for _, annotation := range target.Annotations {
			annotation.Target = target.ID
		}
	}

	result[pkgPath] = targets

	return result, nil
}
