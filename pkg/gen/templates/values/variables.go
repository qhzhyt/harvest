package values

import (
    "github.com/qhzhyt/harvest/pkg/annotation"
)

type Variable struct {
    PackageAlias string
    Name         string
    Target       annotation.Target
}

type VariablesValues struct {
    ImportList   []*Import
    VariableList []*Variable
}

//func NewVariablesValuesMap(annotations []*annotation.Annotation) map[string]*VariablesValues {
//
//    annotationsMap := map[string][]*annotation.Annotation{}
//
//    for _, ann := range annotations {
//        moduleName := strings.ReplaceAll(ann.Target.Module, ".", "_")
//        if _, ok := annotationsMap[moduleName]; !ok {
//            annotationsMap[moduleName] = []*annotation.Annotation{}
//        }
//        annotationsMap[moduleName] = append(annotationsMap[moduleName], ann)
//    }
//
//    instancesValuesMap := map[string]*VariablesValues{}
//
//    for name, anns := range annotationsMap {
//        instancesValuesMap[name] = NewVariablesValues(anns)
//    }
//
//    return instancesValuesMap
//}
//
//func NewVariablesValues(annotations []*annotation.Annotation) *VariablesValues {
//
//    result := &VariablesValues{}
//
//    importMap := map[string]string{}
//
//    for _, ann := range annotations {
//        if ann.Target.Type == annotation.TargetTypeVar && utils.StringFirstIsUpper(ann.Target.Name) {
//            pkgAlias := utils.GenPackageAlias(ann.Target.Package)
//
//            importMap[pkgAlias] = path.Join(ann.Target.Module, ann.Target.Package)
//
//            variable := &Variable{
//                PackageAlias: pkgAlias,
//                Name:         ann.Target.Name,
//                Target:       ann.Target,
//            }
//
//            result.VariableList = append(result.VariableList, variable)
//
//        }
//    }
//
//    for alias, pkg := range importMap {
//        result.ImportList = append(result.ImportList, &Import{
//            Alias:   alias,
//            Package: pkg,
//        })
//    }
//
//    return result
//}
