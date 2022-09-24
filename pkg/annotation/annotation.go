package annotation

type TargetType string

const (
    TargetTypeStruct    TargetType = "struct"
    TargetTypeInterface            = "interface"
    TargetTypeVar                  = "var"
    TargetTypeFunc                 = "func"
    TargetTypeMethod               = "method"
)

type TargetID string

type Annotation struct {
    Name    string
    Args    Args
    Package string
    File    string
    Target  TargetID
}

type Target struct {
    Module      string
    Package     string
    Type        TargetType
    Name        string
    FuncRecv    string
    Annotations []*Annotation
}
