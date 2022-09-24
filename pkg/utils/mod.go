package utils

import (
    "bytes"
    "errors"
    "fmt"
    "io/ioutil"
    "strings"
)

var (
    modulePrefix = []byte("\nmodule ")
)

// GetModulePathFromModFile from: libexec/src/cmd/go/internal/load/pkg.go
func GetModulePathFromModFile(modFilepath string) (string, error) {

    data, err := ioutil.ReadFile(modFilepath)

    if err != nil {
        return "", fmt.Errorf("read mod file failed: %s", err.Error())
    }

    var i int
    if bytes.HasPrefix(data, modulePrefix[1:]) {
        i = 0
    } else {
        i = bytes.Index(data, modulePrefix)
        if i < 0 {
            return "", errors.New("can not found module path in mod file")
        }
        i++
    }
    line := data[i:]

    // Cut line at \n, drop trailing \r if present.
    if j := bytes.IndexByte(line, '\n'); j >= 0 {
        line = line[:j]
    }
    if line[len(line)-1] == '\r' {
        line = line[:len(line)-1]
    }
    line = line[len("module "):]

    // If quoted, unquote.
    return strings.TrimSpace(string(line)), nil
}
