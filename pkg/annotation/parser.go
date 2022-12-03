package annotation

import (
    "fmt"
    "sigs.k8s.io/yaml"
    "strings"
)

func Parse(text string) Annotations {
    var annotations []*Annotation
    var lastAnnotation *Annotation
    for _, a := range strings.Split(text, "\n") {
        a = strings.TrimSpace(a)
        if strings.HasPrefix(a, "@") {
            lastAnnotation = &Annotation{Args: map[string]string{}}
            annotations = append(annotations, lastAnnotation)
            index := strings.Index(a, "(")
            if index > 0 {
                lastAnnotation.Name = Name(a[1:index])
                rightIndex := strings.Index(a, ")")
                if rightIndex > 0 {
                    argText := a[index+1 : rightIndex]
                    for _, arg := range strings.Split(argText, ",") {
                        argSplits := strings.Split(arg, "=")
                        if len(argSplits) > 1 {
                            lastAnnotation.Args[strings.TrimSpace(argSplits[0])] = strings.TrimSpace(argSplits[1])
                        } else {
                            lastAnnotation.Args[strings.TrimSpace(argSplits[0])] = "true"
                        }
                    }
                }
            } else {
                lastAnnotation.Name = Name(a[1:])
            }
            lastAnnotation.Name = Name(strings.ToLower(string(lastAnnotation.Name)))
        } else {
            if lastAnnotation == nil {
                continue
            }
            index := strings.Index(a, ":")
            if index > 0 {
                ss := strings.Split(a, ":")
                lastAnnotation.Args[ss[0]] = ss[1]
            }
        }
    }
    return annotations
}

func Unmarshal(annotation *Annotation, target interface{}) error {
    yamlText := ""
    for k, v := range annotation.Args {
        yamlText += fmt.Sprintf("%s: %s\n", k, v)
    }
    return yaml.Unmarshal([]byte(yamlText), target)
}
