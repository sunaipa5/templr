package main

import (
    "io"
    "os"
    "path/filepath"
)

func copyDir(src, dst string) error {
    return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        rel, _ := filepath.Rel(src, path)
        target := filepath.Join(dst, rel)

        if info.IsDir() {
            return os.MkdirAll(target, info.Mode())
        }

        return copyFile(path, target, info.Mode())
    })
}

func copyFile(src, dst string, mode os.FileMode) error {
    in, err := os.Open(src)
    if err != nil {
        return err
    }
    defer in.Close()

    out, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, mode)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, in)
    return err
}
