package annotation

import (
	"fmt"
	"github.com/qhzhyt/harvest/pkg/utils"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ScanLocalPackage(localPath string, modulePath string, pkg string, deepScan bool) (Targets, error) {
	s := &scanner{
		localPath:  localPath,
		modulePath: modulePath,
		rootPkg:    pkg,
		deepScan:   deepScan,
		targetMap:  map[TargetID]*Target{},
	}
	return s.scan()
}

type scanner struct {
	localPath  string
	modulePath string
	rootPkg    string
	deepScan   bool
	targetMap  map[TargetID]*Target
}

func (s *scanner) scan() (Targets, error) {
	if err := s.scanPkg(s.rootPkg); err != nil {
		return nil, err
	}
	targets := Targets{}
	for _, target := range s.targetMap {
		targets = append(targets, target)
	}
	return targets, nil
}

func (s *scanner) scanPkg(pkg string) error {
	result := map[string]Targets{}

	//var curAnnotations []*Annotation

	var targets []*Target

	fset := token.NewFileSet()
	pkgPath, _ := s.pkgPath(pkg)

	fs, err := parser.ParseDir(fset, pkgPath, nil, parser.ParseComments|parser.AllErrors)

	if err != nil {
		fmt.Println(err)
		return err
	}

	ast.Print(fset, fs)

	for _, f := range fs {

		for fn, file := range f.Files {
			ctx := &fileContext{
				pkg:      pkg,
				pkgName:  file.Name.Name,
				filepath: fn,
				imports:  s.parseImports(file),
			}

			for _, decl := range file.Decls {

				switch d := decl.(type) {
				case *ast.FuncDecl:
					s.parseFuncDecl(ctx, d)
				case *ast.GenDecl:
					s.parseGenDecl(ctx, d)
				default:
					fmt.Println("Decl: ", d)
				}
			}
		}
	}

	if s.deepScan {
		entries, err := os.ReadDir(pkgPath)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			if entry.IsDir() {
				err = s.scanPkg(path.Join(pkg, entry.Name()))
				if err != nil {
					return err
				}
			}
		}
	}

	for _, target := range targets {
		if len(target.ID) == 0 {
			target.ID = GenTargetID(target)
		}
		for _, annotation := range target.Annotations {
			annotation.Target = target.ID
			annotation.ID = GenAnnotationID(target.ID, annotation.Name)
		}
	}

	result[pkgPath] = targets

	return nil
}

func (s *scanner) pkgPath(pkg string) (string, error) {
	pkgPath, err := filepath.Abs(path.Join(s.localPath, pkg))

	if err != nil {
		fmt.Println(err)
	}
	return pkgPath, err
}

func (s *scanner) parseImports(file *ast.File) map[string]string {
	imports := map[string]string{}
	for _, importSpec := range file.Imports {
		var name string
		value := strings.Trim(importSpec.Path.Value, "\"")
		if importSpec.Name == nil || len(importSpec.Name.Name) < 1 {
			_, name = filepath.Split(value)
		} else {
			name = importSpec.Name.Name
		}
		name = strings.Trim(name, "\"")
		imports[name] = value
	}
	return imports
}

func (s *scanner) parseStruct(ctx *fileContext, structType *ast.StructType, target *Target) Targets {
	target.Type = TargetTypeStruct
	target.ID = GenTargetID(target)
	var targetFields []TargetID

	targets := Targets{}

	for _, field := range structType.Fields.List {

		anns := Parse(field.Doc.Text())

		for _, name := range field.Names {

			//switch t := field.Type.(type) {
			//case *ast.Ident:
			//case *ast.StarExpr:
			//case *ast.SelectorExpr:
			//
			//}

			fieldTarget := &Target{
				File:        target.File,
				Module:      target.Module,
				Package:     target.Package,
				Super:       target.ID,
				SuperName:   genSuperName(target.SuperName, string(target.Name)),
				Type:        TargetTypeField,
				Name:        Name(name.Name),
				Annotations: anns,
			}
			fieldTarget.ID = GenTargetID(fieldTarget)
			s.parseField(ctx, field.Type, fieldTarget)
			target.Fields[fieldTarget.Name] = fieldTarget.ID
			targetFields = append(targetFields, fieldTarget.ID)
			targets = append(targets, fieldTarget)
		}

		//switch ft := field.Type.(type) {
		//case *ast.FT:
		//
		//}
		//fmt.Println(field.Tag)
		_ = field
	}

	s.putTargets(targets...)

	return targets
}

func (s *scanner) parseFuncType(ctx *fileContext, funcType *ast.FuncType) *Target {
	target := &Target{
		File:    ctx.filepath,
		Module:  s.modulePath,
		Package: ctx.pkg,
		Fields:  map[Name]TargetID{},
		Type:    TargetTypeFuncType,
	}

	var targets Targets

	for _, field := range s.parseFields(ctx, target, funcType.Params) {
		paramID := TargetID(fmt.Sprintf("%s%s", strings.Repeat("*", int(field.Star)), field.ValueType))
		target.FuncParams = append(target.FuncParams, paramID)
		targets = append(targets, &Target{
			ID:        paramID,
			Name:      field.Name,
			Type:      TargetTypeValueType,
			ValueType: field.ValueType,
			Star:      field.Star,
		})
	}

	for _, field := range s.parseFields(ctx, target, funcType.Results) {
		paramID := TargetID(fmt.Sprintf("%s%s", strings.Repeat("*", int(field.Star)), field.ValueType))
		target.FuncReturns = append(target.FuncReturns, paramID)
		targets = append(targets, &Target{
			ID:        paramID,
			Name:      field.Name,
			Type:      TargetTypeValueType,
			ValueType: field.ValueType,
			Star:      field.Star,
		})
	}

	targetID := fmt.Sprintf("%s:%s", FuncTypeIDPrefix, utils.StringMD5HEXString(
		fmt.Sprintf("(%s):(%s)", strings.Join(target.FuncParams.Strings(), ","), strings.Join(target.FuncReturns.Strings(), ",")),
	))

	target.ID = TargetID(targetID)

	s.putTargets(targets...)
	s.putTargets(target)

	return target
}

func (s *scanner) parseFields(ctx *fileContext, target *Target, fieldList *ast.FieldList) Targets {
	if fieldList == nil {
		return nil
	}
	var targets Targets
	for _, field := range fieldList.List {
		anns := Parse(field.Doc.Text())
		names := field.Names
		if len(names) < 1 {
			names = append(names, &ast.Ident{Name: "_"})
		}
		for _, name := range names {
			fieldTarget := &Target{
				File:        target.File,
				Module:      target.Module,
				Package:     target.Package,
				Super:       target.ID,
				SuperName:   genSuperName(target.SuperName, string(target.Name)),
				Type:        TargetTypeField,
				Name:        Name(name.Name),
				Annotations: anns,
			}
			fieldTarget.ID = GenTargetID(fieldTarget)
			s.parseField(ctx, field.Type, fieldTarget)
			targets = append(targets, fieldTarget)
		}
	}
	return targets
}

func (s *scanner) parseField(ctx *fileContext, typeExpr ast.Expr, target *Target) {
	fieldTargetType := target.Clone()
	fieldTargetType.Fields = map[Name]TargetID{}
	s.parseTypeExpr(ctx, typeExpr, fieldTargetType)
	target.Star = fieldTargetType.Star
	switch fieldTargetType.Type {
	case TargetTypeAlias:
		target.ValueType = fieldTargetType.Super
	case TargetTypeStruct, TargetTypeInterface, TargetTypeFuncType:
		target.ValueType = fieldTargetType.ID
		s.putTargets(fieldTargetType)
	}
}

func (s *scanner) parseFuncDecl(ctx *fileContext, funcDecl *ast.FuncDecl) Targets {
	target := &Target{
		ID:          s.genTargetID(ctx.pkg, TargetTypeFunc, funcDecl.Name.Name),
		File:        ctx.filepath,
		Module:      s.modulePath,
		Package:     ctx.pkg,
		Type:        TargetTypeFunc,
		Name:        Name(funcDecl.Name.Name),
		Annotations: Parse(funcDecl.Doc.Text()),
	}

	//d.Type.Params.List[0]

	if funcDecl.Recv != nil {

		var recvName string

		switch rt := funcDecl.Recv.List[0].Type.(type) {
		case *ast.StarExpr:
			recvName = rt.X.(*ast.Ident).Name
			target.FuncRecv = s.genTypeSpecTargetID(ctx.pkg, rt.X.(*ast.Ident).Obj.Decl.(*ast.TypeSpec))
			target.Star++
		case *ast.Ident:
			recvName = rt.Name
			target.FuncRecv = s.genTypeSpecTargetID(ctx.pkg, rt.Obj.Decl.(*ast.TypeSpec))
		default:
			fmt.Println(rt)
		}
		target.Type = TargetTypeMethod
		target.SuperName = recvName
		target.Super = target.FuncRecv
		target.ID = GenTargetID(target)
		//fmt.Println(a.FuncRecv, a.Name)
	}
	funcType := s.parseFuncType(ctx, funcDecl.Type)
	target.ValueType = funcType.ID
	s.putTargets(target)
	return Targets{target}
}

func (s *scanner) parseRecv() {

}

//func (s *scanner) parseAlias(ctx *fileContext, typeExpr ast.Expr, target *Target) {
//	switch t := typeExpr.(type) {
//	case *ast.Ident:
//		target.Type = TargetTypeAlias
//		//fmt.Println("*ast.Ident", t.Name)
//		if t.Obj == nil {
//			target.Super = genBasicTargetID(t.Name)
//		} else {
//			target.Super = s.genTargetID(ctx.pkg, TargetTypeType, t.Name)
//		}
//		target.SuperName = t.Name
//	case *ast.FuncType:
//	//	FuncType
//	//case *ast.:
//	case *ast.SelectorExpr:
//		//fmt.Println("SelectorExpr: ", typeExpr)
//		pkgNameIdent, ok := t.X.(*ast.Ident)
//		if !ok {
//			panic(fmt.Sprintf("unknown typeExpr: %#v", typeExpr))
//		}
//		GenTargetIDByPkgTypeAndName(ctx.imports[pkgNameIdent.Name], TargetTypeType, t.Sel.Name)
//
//	case *ast.StarExpr:
//		//fmt.Println("StarExpr: ", typeSpec.Type)
//		target.Star++
//		s.parseAlias(ctx, t.X, target)
//	default:
//		fmt.Println("TypeSpec: ", typeExpr)
//	}
//}

func (s *scanner) parseTypeSpec(ctx *fileContext, typeSpec *ast.TypeSpec, annotations Annotations) *Target {

	target := &Target{
		ID:          s.genTargetID(ctx.pkg, TargetTypeType, typeSpec.Name.Name),
		File:        ctx.filepath,
		Module:      s.modulePath,
		Package:     ctx.pkg,
		Name:        Name(typeSpec.Name.Name),
		Annotations: append(annotations.Clone(), Parse(typeSpec.Doc.Text())...),
		Fields:      map[Name]TargetID{},
	}

	s.parseTypeExpr(ctx, typeSpec.Type, target)

	s.putTargets(target)

	return target
}

func (s *scanner) parseTypeExpr(ctx *fileContext, typeExpr ast.Expr, target *Target) {
	switch t := typeExpr.(type) {
	case *ast.StructType:
		s.parseStruct(ctx, t, target)
	case *ast.InterfaceType:
		target.Type = TargetTypeInterface
		target.ID = GenTargetID(target)
		for _, m := range t.Methods.List {
			switch mt := m.Type.(type) {
			case *ast.FuncType:
				_ = mt.Params
				_ = mt.Results
			}
		}
	case *ast.Ident:
		target.Type = TargetTypeAlias
		//fmt.Println("*ast.Ident", t.Name)
		if t.Obj == nil {
			target.Super = genBasicTargetID(t.Name)
		} else {
			target.Super = s.genTargetID(ctx.pkg, TargetTypeType, t.Name)
		}
		//target.SuperName = t.Name
	case *ast.FuncType:
		funcType := s.parseFuncType(ctx, t)
		target.Type = TargetTypeAlias
		target.Super = funcType.ID
	case *ast.SelectorExpr:
		//fmt.Println("SelectorExpr: ", typeExpr)
		target.Type = TargetTypeAlias
		pkgNameIdent, ok := t.X.(*ast.Ident)
		if !ok {
			panic(fmt.Sprintf("unknown typeExpr: %#v", typeExpr))
		}
		target.Super = GenTargetIDByPkgTypeAndName(ctx.imports[pkgNameIdent.Name], TargetTypeType, t.Sel.Name)

	case *ast.StarExpr:
		//fmt.Println("StarExpr: ", typeSpec.Type)
		target.Star++
		s.parseTypeExpr(ctx, t.X, target)
	default:
		fmt.Println("TypeExpr: ", typeExpr)
	}
}

func (s *scanner) parseValueSpec(ctx *fileContext, valueSpec *ast.ValueSpec, annotations Annotations) Targets {

	targets := Targets{}

	for _, name := range valueSpec.Names {
		target := &Target{
			File:        ctx.filepath,
			Module:      s.modulePath,
			Package:     ctx.pkg,
			Type:        TargetTypeVar,
			Name:        Name(name.Name),
			Annotations: append(annotations.Clone(), Parse(valueSpec.Doc.Text())...),
		}

		target.ID = GenTargetID(target)
		s.parseField(ctx, valueSpec.Type, target)
		targets = append(targets, target)
	}

	s.putTargets(targets...)

	return targets
}

func (s *scanner) parseGenDecl(ctx *fileContext, genDecl *ast.GenDecl) {

	annotations := Parse(genDecl.Doc.Text())

	for _, spec := range genDecl.Specs {
		switch realSpec := spec.(type) {
		case *ast.TypeSpec:
			s.parseTypeSpec(ctx, realSpec, annotations)
		case *ast.ValueSpec:
			s.parseValueSpec(ctx, realSpec, annotations)
		default:
			fmt.Println("Spec: ", realSpec)
		}
	}
}

func (s *scanner) genTypeSpecTargetID(pkg string, typeSpec *ast.TypeSpec) TargetID {
	switch typeSpec.Type.(type) {
	case *ast.StructType:
		return s.genTargetID(pkg, TargetTypeStruct, typeSpec.Name.Name)
	case *ast.InterfaceType:
		return s.genTargetID(pkg, TargetTypeInterface, typeSpec.Name.Name)

	case *ast.Ident:
		return s.genTargetID(pkg, TargetTypeType, typeSpec.Name.Name)
	default:
		fmt.Println("TypeSpec: ", typeSpec.Type)
		return s.genTargetID(pkg, TargetTypeType, typeSpec.Name.Name)
	}
}

func (s *scanner) genTargetID(pkg string, targetType TargetType, targetName string) TargetID {
	return GenTargetIDByNames(s.modulePath, pkg, targetType, targetName)
}

func (s *scanner) putTargets(targets ...*Target) {
	for _, target := range targets {
		s.targetMap[target.ID] = target
	}
}

func (s *scanner) getTarget(id TargetID) *Target {
	return s.targetMap[id]
}
