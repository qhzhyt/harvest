package main

import (
    "fmt"
    "os"

    "github.com/qhzhyt/harvest/pkg/gen"
)

func main() {
    fmt.Println("harvest")
    //fmt.Println(os.Args)
    //fmt.Println(os.Environ())
    mainFilePath := os.Getenv("GOFILE")
    if len(mainFilePath) == 0 {
        mainFilePath = "main.go"
    }

    harvestGenerator, err := gen.NewGeneratorByMainFilePath(mainFilePath)

    if err != nil {
        panic(err)
    }

    _ = harvestGenerator.GenerateAll()

}
