package annotation

import "fmt"

type TargetID string
type TargetType string
type Star int

const (
	TargetTypeType          TargetType = "type"
	TargetTypeStruct        TargetType = "struct"
	TargetTypeInterface     TargetType = "interface"
	TargetTypeAlias         TargetType = "alias"
	TargetTypeBasicType     TargetType = "basic-type"
	TargetTypeStructType    TargetType = "struct-type"
	TargetTypeInterfaceType TargetType = "interface-type"
	TargetTypeFuncType      TargetType = "func-type"
	TargetTypeValueType     TargetType = "value-type"
	TargetTypeVar           TargetType = "var"
	TargetTypeFunc          TargetType = "func"
	TargetTypeMethod        TargetType = "method"
	TargetTypeField         TargetType = "field"
)

const (
	FuncTypeIDPrefix = "--func--"
	BasicTypePrefix  = "--basic--"
)

type Target struct {
	ID          TargetID
	File        string
	Module      string
	Package     string
	Type        TargetType
	Name        Name
	Super       TargetID
	SuperName   string
	ValueType   TargetID
	Methods     TargetIDMap
	Fields      TargetIDMap
	FuncRecv    TargetID
	Star        Star
	FuncParams  TargetIDList
	FuncReturns TargetIDList
	Annotations Annotations
}

func (t *Target) Clone() *Target {
	return &Target{
		ID:          t.ID,
		File:        t.File,
		Module:      t.Module,
		Package:     t.Package,
		Type:        t.Type,
		Name:        t.Name,
		Super:       t.Super,
		SuperName:   t.SuperName,
		ValueType:   t.ValueType,
		Methods:     t.Methods,
		Fields:      t.Fields,
		FuncRecv:    t.FuncRecv,
		Star:        t.Star,
		FuncParams:  t.FuncParams,
		FuncReturns: t.FuncReturns,
		Annotations: t.Annotations,
	}
}

func (tid TargetID) Get() *Target {
	return GetTargetByID(tid)
}

func genBasicTargetID(typeName string) TargetID {
	return TargetID(fmt.Sprintf("%s:%s", BasicTypePrefix, typeName))
}

func genSuperName(superSuperName, superName string) string {
	if len(superSuperName) < 1 {
		return superName
	}
	return superSuperName + "." + superName
}

func GenTargetID(target *Target) TargetID {
	return GenTargetIDByNames(target.Module, target.Package, target.Type, genSuperName(target.SuperName, string(target.Name)))
}

func GenTargetIDByNames(module string, pkg string, targetType TargetType, targetName string) TargetID {
	switch targetType {
	case TargetTypeStruct, TargetTypeInterface:
		targetType = TargetTypeType
	}

	return TargetID(fmt.Sprintf("%s/%s:%s:%s", module, pkg, targetType, targetName))
}

func GenTargetIDByPkgTypeAndName(pkg string, targetType TargetType, targetName string) TargetID {
	switch targetType {
	case TargetTypeStruct, TargetTypeInterface:
		targetType = TargetTypeType
	}
	return TargetID(fmt.Sprintf("%s:%s:%s", pkg, targetType, targetName))
}

type Targets []*Target
type TargetIDMap map[Name]TargetID
type TargetIDList []TargetID

func (t TargetIDList) Strings() []string {
	strs := make([]string, len(t))
	for i, id := range t {
		strs[i] = string(id)
	}
	return strs
}
