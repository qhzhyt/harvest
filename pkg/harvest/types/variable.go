package types

import (
    "github.com/qhzhyt/harvest/pkg/annotation"
    "reflect"
)

type Variable struct {
    annotation.Target
    PtrValue   reflect.Value
    Annotation *annotation.Annotation
}
