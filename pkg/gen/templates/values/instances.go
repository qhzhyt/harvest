package values

import (
    "github.com/qhzhyt/harvest/pkg/annotation"
    "github.com/qhzhyt/harvest/pkg/utils"
    "path"
    "strings"
)

type Import struct {
    Alias   string
    Package string
}

type Instance struct {
    PackageAlias    string
    TypeName        string
    InstanceName    string
    InstancePtrName string
}

type Component struct {
    annotation.Target
    InstanceName string
}

type InstancesValues struct {
    ImportList    []*Import
    InstanceList  []*Instance
    ComponentList []*Component
}

func NewInstancesValuesMap(annotations []*annotation.Annotation) map[string]*InstancesValues {

    annotationsMap := map[string][]*annotation.Annotation{}

    for _, ann := range annotations {
        moduleName := strings.ReplaceAll(ann.Target.Module, ".", "_")
        if _, ok := annotationsMap[moduleName]; !ok {
            annotationsMap[moduleName] = []*annotation.Annotation{}
        }
        annotationsMap[moduleName] = append(annotationsMap[moduleName], ann)
    }

    instancesValuesMap := map[string]*InstancesValues{}

    for name, anns := range annotationsMap {
        instancesValuesMap[name] = NewInstancesValues(anns)
    }

    return instancesValuesMap
}

func NewInstancesValues(annotations []*annotation.Annotation) *InstancesValues {

    result := &InstancesValues{}

    importMap := map[string]string{}

    for _, ann := range annotations {
        if ann.Target.Type == annotation.TargetTypeStruct && utils.StringFirstIsUpper(ann.Target.Name) {
            pkgAlias := utils.GenPackageAlias(ann.Target.Package)

            importMap[pkgAlias] = path.Join(ann.Target.Module, ann.Target.Package)

            instance := &Instance{
                PackageAlias:    pkgAlias,
                TypeName:        ann.Target.Name,
                InstanceName:    utils.StringFirstLower(ann.Target.Name) + "Instance",
                InstancePtrName: utils.StringFirstLower(ann.Target.Name) + "Ptr",
            }

            result.InstanceList = append(result.InstanceList, instance)

            result.ComponentList = append(result.ComponentList, &Component{
                Target:       ann.Target,
                InstanceName: instance.InstancePtrName,
            })

        }
    }

    for alias, pkg := range importMap {
        result.ImportList = append(result.ImportList, &Import{
            Alias:   alias,
            Package: pkg,
        })
    }

    return result
}
