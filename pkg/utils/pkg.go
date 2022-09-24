package utils

import "strings"

// GenPackageAlias 根据包路径生成包别名
func GenPackageAlias(pkg string) string {
    pkgSplit := strings.Split(pkg, "/")
    pkgAlias := ""
    for _, n := range pkgSplit {
        pkgAlias += StringFirstUpper(n)
    }
    return StringFirstLower(pkgAlias)
}
