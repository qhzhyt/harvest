package harvest

import (
    "reflect"
)

import (
    components "simple/components"
)

import (
    "github.com/qhzhyt/harvest/pkg/annotation"
    "github.com/qhzhyt/harvest/pkg/harvest/types"
)

var Variables = []*types.Variable{
    {
        Target: annotation.Target{
            Module:   "simple",
            Package:  "components",
            Type:     "struct",
            Name:     "Collection",
            FuncRecv: "",
        },
        PtrValue: reflect.ValueOf(&components.UserServiceImpl).Elem(),
    },
}
