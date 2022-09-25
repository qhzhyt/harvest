package gen

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/qhzhyt/harvest/pkg/annotation"
	"github.com/qhzhyt/harvest/pkg/harvest"
	"github.com/qhzhyt/harvest/pkg/harvest/types"
	"github.com/qhzhyt/harvest/pkg/log"
	"github.com/qhzhyt/harvest/pkg/utils"
)

type HarvestGenerator struct {
	*types.Harvest
}

func (g *HarvestGenerator) processTargets(targets []*annotation.Target) {
	for _, target := range targets {
		if strings.HasPrefix(target.Module, "/") || utils.IsDir(target.Module) {
			target.Module = g.ModulePath
		}
	}
}

func (g *HarvestGenerator) GenerateAll() error {
	targetsMap := map[string][]*annotation.Target{}

	for _, pkg := range g.ScanPath {

		//if !strings.HasPrefix(pkg, "/") {
		//    pkg = path.Join(path.Dir(g.MainFilePath), pkg)
		//}

		targets, err := annotation.ScanLocalPackage(g.ModuleRoot, g.ModulePath, pkg, g.DeepScan)
		if err != nil {
			panic(err)
		}

		for pkg0, tgs := range targets {
			targetsMap[pkg0] = tgs
		}
	}

	var targets []*annotation.Target

	for _, tgs := range targetsMap {
		targets = append(targets, tgs...)
	}

	g.processTargets(targets)

	bytes, _ := json.MarshalIndent(targets, " ", "  ")
	fmt.Println(string(bytes))

	if exists, err := utils.PathExists(g.OutputPath); !exists {
		if err != nil {
			log.ErrorAndPanic(err)
		}
		err = os.MkdirAll(g.OutputPath, 0777)
		if err != nil {
			log.ErrorAndPanic(err)
		}
	} else {
		if !utils.IsDir(g.OutputPath) {
			log.ErrorAndExit(g.OutputPath, "is exists, but not a dir")
		}
	}

	//err := templates.ExecuteTemplateToFile(templates.AnnotationsTemplate, values.AnnotationsValues{Annotations: annotations}, path.Join(g.OutputPath, string(templates.AnnotationsTemplate)))
	//
	//if err != nil {
	//    log.ErrorAndPanic(err)
	//}
	//
	//instancesValuesMap := values.NewInstancesValuesMap(annotations)
	//
	//for name, instancesValues := range instancesValuesMap {
	//    log.Println("------", "instances_"+name+".go", "------")
	//    err = templates.ExecuteTemplateToFile(templates.InstancesTemplate, instancesValues, path.Join(g.OutputPath, "instances_"+name+".go"))
	//    if err != nil {
	//        log.ErrorAndPanic(err)
	//    }
	//}
	//
	//variablesValuesMap := values.NewVariablesValuesMap(annotations)
	//
	//for name, variablesValues := range variablesValuesMap {
	//    log.Println("------", "variables_"+name+".go", "------")
	//    err = templates.ExecuteTemplateToFile(templates.VariablesTemplate, variablesValues, path.Join(g.OutputPath, "variables_"+name+".go"))
	//    if err != nil {
	//        log.ErrorAndPanic(err)
	//    }
	//}

	return nil
}

func NewGenerator(harvestConfig *types.Harvest) *HarvestGenerator {
	return &HarvestGenerator{Harvest: harvestConfig}
}

func NewGeneratorByMainFilePath(mainFilePath string) (*HarvestGenerator, error) {

	harvestConfig, err := harvest.ParseMain(mainFilePath)

	if err != nil {
		return nil, err
	}

	harvestConfig.MainFilePath = mainFilePath

	if len(harvestConfig.ModulePath) < 1 {
		harvestConfig.ModulePath, err = utils.GetModulePathFromModFile(path.Join(path.Dir(mainFilePath), "go.mod"))
		if err != nil {
			return nil, err
		}
	}

	harvestConfig.ModuleRoot, err = filepath.Abs(path.Dir(mainFilePath))

	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(harvestConfig.OutputPath, "/") {
		harvestConfig.OutputPath = path.Join(harvestConfig.ModuleRoot, harvestConfig.OutputPath)
	}

	fmt.Println(harvestConfig)

	return NewGenerator(harvestConfig), nil
}
