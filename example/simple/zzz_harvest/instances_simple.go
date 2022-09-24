package harvest

import (
    components "simple/components"
)

import (
    "github.com/qhzhyt/harvest/pkg/annotation"
    "github.com/qhzhyt/harvest/pkg/harvest/types"
)

var (
    collectionInstance = components.Collection{}
    userInstance       = components.User{}
)

var (
    collectionPtr = &collectionInstance
    userPtr       = &userInstance
)

var Components = []*types.Component{
    {
        Target: annotation.Target{
            Module:   "simple",
            Package:  "components",
            Type:     "struct",
            Name:     "Collection",
            FuncRecv: "",
        },
        Instance: collectionPtr,
    },
    {
        Target: annotation.Target{
            Module:   "simple",
            Package:  "components",
            Type:     "struct",
            Name:     "Collection",
            FuncRecv: "",
        },
        Instance: userPtr,
    },
}
