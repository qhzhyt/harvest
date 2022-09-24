package utils

import "os"

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
    s, err := os.Stat(path)
    if err != nil {

        return false
    }
    return s.IsDir()
}

// PathExists 判断所给路径文件/文件夹是否存在
func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}
