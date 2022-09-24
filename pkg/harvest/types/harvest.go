package types

type Harvest struct {
    ScanPath     []string `yaml:"scanPath"`
    DeepScan     bool     `yaml:"deepScan"`
    OutputPath   string   `yaml:"outputPath"`
    ModulePath   string   `yaml:"modulePath"`
    MainFilePath string   `yaml:"-"`
    ModuleRoot   string   `yaml:"-"`
}
