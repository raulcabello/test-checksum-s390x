package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"syscall"
)

func main() {
	str, err := getFileChecksum("test")
	if err != nil {
		fmt.Println("error " + err.Error())
	} else {
		fmt.Println("checksum" + str)
	}
}

func getFileChecksum(filePath string) (string, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY|syscall.O_DIRECT, 0)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))
	return hex.EncodeToString(h.Sum(nil)), nil
}
