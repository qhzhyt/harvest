package main

import (
    "fmt"
    "reflect"
    "simple/components"

    _ "simple/zzz_harvest"
)

var user = &components.User{}

//go:generate harvest-gen
// @harvest(autoConfig, deepScan)
// scanPath: [components]
// outputPath: zzz_harvest
func main() {
    fmt.Println("simple")

    u := reflect.ValueOf(user).Elem()

    u.Set(reflect.ValueOf(components.User{
        ID:         6,
        Name:       "name",
        Collection: nil,
    }))

    fmt.Println(user)

}
