package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	PrintModes("SmartFind_0A33908.mhtml")
	PrintModes("SmartFind_4XH0L55146.mhtml")
}

func PrintModes(filePath string) {
	fileContent, _ := ReadTxtToString(filePath)

	//只取 tbody 區塊
	fisrtTbodyIdx := strings.Index(fileContent, "<tbody>")
	lastTbodyIdx := strings.LastIndex(fileContent, "</tbody>")
	fileContent = fileContent[fisrtTbodyIdx:lastTbodyIdx]

	//依照關鍵字開頭分割
	splitedFileContent := strings.Split(fileContent, "ThinkPad")

	for _, txtline := range splitedFileContent {
		cutIdx := strings.IndexFunc(txtline, func(r rune) bool { return '=' == r || '<' == r })
		if -1 != cutIdx {
			fmt.Printf("ThinkPad%s\n", txtline[:cutIdx])
		} else {
			fmt.Printf("ThinkPad%s\n", txtline)
		}
	}
}

func ReadTxtToString(filename string) (string, error) {
	var content string

	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	for {
		b := make([]byte, 100)
		count, err := file.Read(b)
		if err != nil {
			if io.EOF == err {
				break
			} else {
				log.Fatal(err)
				return "", err
			}
		}
		// fmt.Printf("read %d bytes\n", count)
		// fmt.Printf("last byte :%v\n", b[count-1])
		content += string(b[:count])
	}

	return content, nil
}
