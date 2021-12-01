package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/longhorn/sparse-tools/sparse"
	"io"
	"os"
)

func main() {
	str, err := GetFileChecksum("test")
	if err != nil {
		fmt.Println("error " + err.Error())
	} else {
		fmt.Println(str)
	}
}

func GetFileChecksum(filePath string) (string, error) {
	f, err := sparse.NewDirectFileIoProcessor(filePath, os.O_RDONLY, 0)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// 4MB
	buf := make([]byte, 1<<22)
	h := sha512.New()

	for {
		nr, err := f.Read(buf)
		if err != nil {
			if err != io.EOF {
				return "", err
			}
			break
		}
		h.Write(buf[:nr])
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
