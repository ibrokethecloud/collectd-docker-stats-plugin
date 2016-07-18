package main 

import (
	//"fmt"
	"os"
	"path/filepath"
)

func DirSizeMB(path string) (int64, error) {
    var size int64
    err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            size += info.Size()
        }
        return err
    })
    size = size/(1024*1024)
    return size, err
}