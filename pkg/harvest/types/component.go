package types

import (
    "github.com/qhzhyt/harvest/pkg/annotation"
)

type Component struct {
    Target     annotation.TargetID
    Instance   interface{}
    Annotation *annotation.Annotation
}
