package harvest

import "github.com/qhzhyt/harvest/pkg/annotation"

var Targets = []*annotation.Target {
    {
        ID:             "--basic--:bool",
        File:           "",
        Module:         "",
        Package:        "",
        Type:           "value-type",
        Name:           "_",
        Super:          "",
        SuperName:      "",
        ValueType:      "--basic--:bool",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:field:User.Collection",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "field",
        Name:           "Collection",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        SuperName:      "User",
        ValueType:      "--basic--:Collection",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           1,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:field:User.Service",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "field",
        Name:           "Service",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        SuperName:      "User",
        ValueType:      "github.com/qhzhyt/harvest/examples/simple/components:type:UserService",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "service",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "--func--:032c50e2cb9507ff87a4a1556ae29821",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "func-type",
        Name:           "",
        Super:          "",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
            "--basic--:string",
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:method:User.SetName",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "method",
        Name:           "SetName",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        SuperName:      "User",
        ValueType:      "--func--:032c50e2cb9507ff87a4a1556ae29821",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           1,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "component",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:func:String",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "func",
        Name:           "String",
        Super:          "",
        SuperName:      "",
        ValueType:      "--func--:ff2dfe92f11535551368428d03fbdbd1",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "func",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:field:User.Label",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "field",
        Name:           "Label",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        SuperName:      "User",
        ValueType:      "--basic--:string",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:field:NameStruct0.Name",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "field",
        Name:           "Name",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:NameStruct0",
        SuperName:      "NameStruct0",
        ValueType:      "--basic--:string",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components/sub:type:Sub",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/sub/sub.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components/sub",
        Type:           "struct",
        Name:           "Sub",
        Super:          "",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "component",
                Args:    map[string]string {
                },
                Target:  "",
            },
            {
                ID:      "",
                Name:    "sub",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "*github.com/qhzhyt/harvest/examples/simple/components:type:User",
        File:           "",
        Module:         "",
        Package:        "",
        Type:           "value-type",
        Name:           "_",
        Super:          "",
        SuperName:      "",
        ValueType:      "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           1,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:field:NameStruct.Name",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "field",
        Name:           "Name",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:NameStruct",
        SuperName:      "NameStruct",
        ValueType:      "--basic--:string",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:method:int0.String",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "method",
        Name:           "String",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:int0",
        SuperName:      "int0",
        ValueType:      "--func--:ff2dfe92f11535551368428d03fbdbd1",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "github.com/qhzhyt/harvest/examples/simple/components:type:int0",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           1,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "method",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:func:NewUser",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "func",
        Name:           "NewUser",
        Super:          "",
        SuperName:      "",
        ValueType:      "--func--:2c8a4210091b17bee45913f02762ea87",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "constructor",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components/sub:func:test",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/sub/sub.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components/sub",
        Type:           "func",
        Name:           "test",
        Super:          "",
        SuperName:      "",
        ValueType:      "--func--:aa9d9c180f39535d09a36e6aa24f2420",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "test",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "--func--:623f0d3ad907f0f2dc08abe6dd91f972",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "func-type",
        Name:           "",
        Super:          "",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
            "*github.com/qhzhyt/harvest/examples/simple/components:type:User",
        },
        FuncReturns:    annotation.TargetIDList {
            "*github.com/qhzhyt/harvest/examples/simple/components:type:User",
            "--basic--:error",
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "struct",
        Name:           "User",
        Super:          "",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
            "Collection": "github.com/qhzhyt/harvest/examples/simple/components:field:User.Collection",
            "Handler": "github.com/qhzhyt/harvest/examples/simple/components:field:User.Handler",
            "HandlerFunc": "github.com/qhzhyt/harvest/examples/simple/components:field:User.HandlerFunc",
            "ID": "github.com/qhzhyt/harvest/examples/simple/components:field:User.ID",
            "IP": "github.com/qhzhyt/harvest/examples/simple/components:field:User.IP",
            "Label": "github.com/qhzhyt/harvest/examples/simple/components:field:User.Label",
            "Name": "github.com/qhzhyt/harvest/examples/simple/components:field:User.Name",
            "Service": "github.com/qhzhyt/harvest/examples/simple/components:field:User.Service",
            "Struct": "github.com/qhzhyt/harvest/examples/simple/components:field:User.Struct",
            "Sub": "github.com/qhzhyt/harvest/examples/simple/components:field:User.Sub",
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "component",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:type:int0",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "alias",
        Name:           "int0",
        Super:          "--basic--:int",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "component",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "*github.com/qhzhyt/harvest/examples/simple/components/sub:type:Sub",
        File:           "",
        Module:         "",
        Package:        "",
        Type:           "value-type",
        Name:           "sub",
        Super:          "",
        SuperName:      "",
        ValueType:      "github.com/qhzhyt/harvest/examples/simple/components/sub:type:Sub",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           1,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:type:User.Struct",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "struct",
        Name:           "Struct",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        SuperName:      "User",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
            "Name": "github.com/qhzhyt/harvest/examples/simple/components:field:User.Struct.Name",
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "struct",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:field:User.Sub",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "field",
        Name:           "Sub",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        SuperName:      "User",
        ValueType:      ":type:Sub",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "sub",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:field:User.Handler",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "field",
        Name:           "Handler",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        SuperName:      "User",
        ValueType:      "github.com/qhzhyt/harvest/examples/simple/components:type:User.Handler",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "interface",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:type:uStruct",
        File:           "",
        Module:         "",
        Package:        "",
        Type:           "value-type",
        Name:           "uStruct",
        Super:          "",
        SuperName:      "",
        ValueType:      "github.com/qhzhyt/harvest/examples/simple/components:type:uStruct",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:var:UserServiceImpl",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "var",
        Name:           "UserServiceImpl",
        Super:          "",
        SuperName:      "",
        ValueType:      "github.com/qhzhyt/harvest/examples/simple/components:type:UserService",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "component",
                Args:    map[string]string {
                    "name": "userService",
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:type:User.Handler",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "interface",
        Name:           "Handler",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        SuperName:      "User",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "interface",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "--basic--:string",
        File:           "",
        Module:         "",
        Package:        "",
        Type:           "value-type",
        Name:           "_",
        Super:          "",
        SuperName:      "",
        ValueType:      "--basic--:string",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:method:User.GetName",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "method",
        Name:           "GetName",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        SuperName:      "User",
        ValueType:      "--func--:ff2dfe92f11535551368428d03fbdbd1",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "component",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "--func--:aa9d9c180f39535d09a36e6aa24f2420",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/sub/sub.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components/sub",
        Type:           "func-type",
        Name:           "",
        Super:          "",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:field:uStruct.Name",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "field",
        Name:           "Name",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:uStruct",
        SuperName:      "uStruct",
        ValueType:      "--basic--:string",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components/sub:type:number",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/sub/sub.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components/sub",
        Type:           "alias",
        Name:           "number",
        Super:          "--basic--:int64",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "component",
                Args:    map[string]string {
                },
                Target:  "",
            },
            {
                ID:      "",
                Name:    "type",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components/sub:method:add.test",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/sub/sub.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components/sub",
        Type:           "method",
        Name:           "test",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components/sub:type:add",
        SuperName:      "add",
        ValueType:      "--func--:aa9d9c180f39535d09a36e6aa24f2420",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "github.com/qhzhyt/harvest/examples/simple/components/sub:type:add",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           1,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:field:User.ID",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "field",
        Name:           "ID",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        SuperName:      "User",
        ValueType:      "--basic--:int64",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "auto",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:field:User.Struct",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "field",
        Name:           "Struct",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        SuperName:      "User",
        ValueType:      "github.com/qhzhyt/harvest/examples/simple/components:type:User.Struct",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "struct",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components/sub:type:netIP",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/sub/sub.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components/sub",
        Type:           "alias",
        Name:           "netIP",
        Super:          "net:type:IP",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "component",
                Args:    map[string]string {
                },
                Target:  "",
            },
            {
                ID:      "",
                Name:    "type",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:type:NameStruct0",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "struct",
        Name:           "NameStruct0",
        Super:          "",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
            "Name": "github.com/qhzhyt/harvest/examples/simple/components:field:NameStruct0.Name",
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:type:SuperUser",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "alias",
        Name:           "SuperUser",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "component",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "--func--:2c8a4210091b17bee45913f02762ea87",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "func-type",
        Name:           "",
        Super:          "",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
            "github.com/qhzhyt/harvest/examples/simple/components:type:uStruct",
        },
        FuncReturns:    annotation.TargetIDList {
            "*github.com/qhzhyt/harvest/examples/simple/components:type:User",
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components/sub:type:add",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/sub/sub.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components/sub",
        Type:           "struct",
        Name:           "add",
        Super:          "",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "component",
                Args:    map[string]string {
                },
                Target:  "",
            },
            {
                ID:      "",
                Name:    "add",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components/sub:type:starSub",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/sub/sub.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components/sub",
        Type:           "alias",
        Name:           "starSub",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components/sub:type:Sub",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           1,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "component",
                Args:    map[string]string {
                },
                Target:  "",
            },
            {
                ID:      "",
                Name:    "type",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components/sub:type:subSub",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/sub/sub.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components/sub",
        Type:           "alias",
        Name:           "subSub",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components/sub:type:Sub",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "component",
                Args:    map[string]string {
                },
                Target:  "",
            },
            {
                ID:      "",
                Name:    "type",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:type:UserHandlerFunc",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "alias",
        Name:           "UserHandlerFunc",
        Super:          "--func--:623f0d3ad907f0f2dc08abe6dd91f972",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "type",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components/sub:type:handlerFunc",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/sub/sub.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components/sub",
        Type:           "alias",
        Name:           "handlerFunc",
        Super:          "--func--:073674a114d6b530a5b34a8fefd4580a",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "component",
                Args:    map[string]string {
                },
                Target:  "",
            },
            {
                ID:      "",
                Name:    "type",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:field:User.IP",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "field",
        Name:           "IP",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        SuperName:      "User",
        ValueType:      "net:type:IP",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "ip",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:type:NameStruct",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "struct",
        Name:           "NameStruct",
        Super:          "",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
            "Name": "github.com/qhzhyt/harvest/examples/simple/components:field:NameStruct.Name",
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:type:Collection",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/collection.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "struct",
        Name:           "Collection",
        Super:          "",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
            "ID": "github.com/qhzhyt/harvest/examples/simple/components:field:Collection.ID",
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "component",
                Args:    map[string]string {
                },
                Target:  "",
            },
            {
                ID:      "",
                Name:    "controller",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:field:User.HandlerFunc",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "field",
        Name:           "HandlerFunc",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        SuperName:      "User",
        ValueType:      "--func--:623f0d3ad907f0f2dc08abe6dd91f972",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "func",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components/sub:type:starStarSubSub",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/sub/sub.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components/sub",
        Type:           "alias",
        Name:           "starStarSubSub",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components/sub:type:subSub",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           2,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "component",
                Args:    map[string]string {
                },
                Target:  "",
            },
            {
                ID:      "",
                Name:    "type",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:field:Collection.ID",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/collection.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "field",
        Name:           "ID",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:Collection",
        SuperName:      "Collection",
        ValueType:      "--basic--:int64",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "--basic--:error",
        File:           "",
        Module:         "",
        Package:        "",
        Type:           "value-type",
        Name:           "_",
        Super:          "",
        SuperName:      "",
        ValueType:      "--basic--:error",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:field:User.Struct.Name",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "field",
        Name:           "Name",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:User.Struct",
        SuperName:      "User.Struct",
        ValueType:      "--basic--:string",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:field:User.Name",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "field",
        Name:           "Name",
        Super:          "github.com/qhzhyt/harvest/examples/simple/components:type:User",
        SuperName:      "User",
        ValueType:      "--basic--:string",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components:type:UserService",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "interface",
        Name:           "UserService",
        Super:          "",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "--func--:ff2dfe92f11535551368428d03fbdbd1",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/user.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components",
        Type:           "func-type",
        Name:           "",
        Super:          "",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
            "--basic--:string",
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "--func--:073674a114d6b530a5b34a8fefd4580a",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/sub/sub.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components/sub",
        Type:           "func-type",
        Name:           "",
        Super:          "",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
            "*github.com/qhzhyt/harvest/examples/simple/components/sub:type:Sub",
        },
        FuncReturns:    annotation.TargetIDList {
            "--basic--:bool",
        },
        Star:           0,
        Annotations:    annotation.Annotations{
        },
    },
    {
        ID:             "github.com/qhzhyt/harvest/examples/simple/components/sub:type:starIP",
        File:           "/Users/heyitong/Projects/GolandProjects/harvest/examples/simple/components/sub/sub.go",
        Module:         "github.com/qhzhyt/harvest/examples/simple",
        Package:        "components/sub",
        Type:           "alias",
        Name:           "starIP",
        Super:          "net:type:IP",
        SuperName:      "",
        ValueType:      "",
        Methods:        map[annotation.Name]annotation.TargetID {
        },
        Fields:         map[annotation.Name]annotation.TargetID {
        },
        FuncRecv:       "",
        FuncParams:     annotation.TargetIDList {
        },
        FuncReturns:    annotation.TargetIDList {
        },
        Star:           1,
        Annotations:    annotation.Annotations{
            {
                ID:      "",
                Name:    "component",
                Args:    map[string]string {
                },
                Target:  "",
            },
            {
                ID:      "",
                Name:    "type",
                Args:    map[string]string {
                },
                Target:  "",
            },
        },
    },
}
