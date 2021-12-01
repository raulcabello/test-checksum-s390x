package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
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
	//f, err := sparse.NewDirectFileIoProcessor(filePath, os.O_RDONLY, 0)
	//if err != nil {
	//	return "", err
	//}
	//defer f.Close()

	// 4MB
	//buf := make([]byte, 1<<22)
	//h := sha512.New()
	//
	//for {
	//	nr, err := f.Read(buf)
	//	if err != nil {
	//		if err != io.EOF {
	//			return "", err
	//		}
	//		break
	//	}
	//	h.Write(buf[:nr])
	//}
	f, err := os.Open(filePath)
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
