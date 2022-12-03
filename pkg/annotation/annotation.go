package annotation

import "fmt"

type Name string
type ID string
type Args map[string]string

type Annotation struct {
    ID     ID
    Name   Name
    Args   Args
    Target TargetID
}

func (a *Annotation) Clone() *Annotation {
    return &Annotation{
        ID:     a.ID,
        Name:   a.Name,
        Args:   a.Args,
        Target: a.Target,
    }
}

type Annotations []*Annotation

func (as Annotations) Clone() Annotations {
    result := Annotations{}
    for _, a := range as {
        result = append(result, a.Clone())
    }
    return result
}

func (id ID) Get() *Annotation {
    return GetAnnotationByID(id)
}

func GenAnnotationID(targetID TargetID, annName Name) ID {
    return ID(fmt.Sprintf("%s:%s", targetID, annName))
}
