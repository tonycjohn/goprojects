package main

import (
	"io"
	"log"
	"os"
)

func copyFile(dstFile string, srcFile string) (written int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println("Bad Error:", err)
		}
	}()

	defer src.Close()

	dst, err := os.Create(dstFile)
	if err != nil {
		return
	}
	defer dst.Close()

	written, err = io.Copy(dst, src)
	return
}
