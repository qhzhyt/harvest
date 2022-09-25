package annotation

import "fmt"

type TargetType string

const (
	TargetTypeStruct    TargetType = "struct"
	TargetTypeInterface            = "interface"
	TargetTypeVar                  = "var"
	TargetTypeFunc                 = "func"
	TargetTypeMethod               = "method"
)

type TargetID string

type Args map[string]string

type Annotation struct {
	Name    string
	Args    Args
	Package string
	File    string
	Target  TargetID
}

type Target struct {
	ID          TargetID
	File        string
	Module      string
	Package     string
	Type        TargetType
	Name        string
	FuncRecv    string
	Annotations []*Annotation
}

func (t *Target) Clone() *Target {
	return &Target{
		ID:          t.ID,
		File:        t.File,
		Module:      t.Module,
		Package:     t.Package,
		Type:        t.Type,
		Name:        t.Name,
		FuncRecv:    t.FuncRecv,
		Annotations: t.Annotations,
	}
}

func GenTargetID(target *Target) TargetID {
	return TargetID(fmt.Sprintf("%s:%s:%s:%s", target.Module, target.Package, target.Type, target.Name))
}
