package annotation

import "github.com/qhzhyt/harvest/pkg/utils"

var basicTypeNames = []Name{
	"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64",
	"byte", "rune", "bool", "float32", "float64", "string",
}

func GetBasicValueType() Targets {
	return utils.Map[Name, *Target](basicTypeNames, func(name Name) *Target {
		return &Target{
			ID:   genBasicTargetID(string(name)),
			Type: TargetTypeBasicType,
			Name: name,
		}
	})
}
