package templates

import (
    "embed"
    "fmt"
    "os"
    "strings"
    "text/template"

    "github.com/Masterminds/sprig/v3"

    "github.com/qhzhyt/harvest/pkg/utils"
)

//go:embed _templates
var templates embed.FS

var templateAll *template.Template

func init() {

    funcMap := sprig.TxtFuncMap()

    funcMap["include"] = func(name string, data interface{}) (string, error) {
        var buf strings.Builder
        err := templateAll.ExecuteTemplate(&buf, name, data)
        return buf.String(), err
    }

    var err error
    templateAll, err = template.New("harvest").Funcs(funcMap).ParseFS(templates, "_templates/common/*", "_templates/*.tmpl")
    if err != nil {
        panic(fmt.Sprintf("templates parse failed: %s", err.Error()))
    }
}

func ExecuteTemplate(name TemplateName, input interface{}) ([]byte, error) {
    output := &utils.TempWriter{}
    err := templateAll.ExecuteTemplate(output, string(name), input)
    if err != nil {
        return nil, err
    }
    return output.Bytes(), nil
}

func ExecuteTemplateToFile(name TemplateName, input interface{}, filename string) error {

    fp, err := os.Create(filename)

    if err != nil {
        return err
    }

    err = templateAll.ExecuteTemplate(fp, string(name), input)
    if err != nil {
        return err
    }
    return nil
}
