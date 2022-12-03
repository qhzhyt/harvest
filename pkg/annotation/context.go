package annotation

type packageContext struct {
	packageName string
}

type fileContext struct {
	pkg      string
	pkgName  string
	filepath string
	imports  map[string]string
}
