package types

import (
    "github.com/qhzhyt/harvest/pkg/annotation"
)

type Component struct {
    annotation.Target
    Instance   interface{}
    Annotation *annotation.Annotation
}
