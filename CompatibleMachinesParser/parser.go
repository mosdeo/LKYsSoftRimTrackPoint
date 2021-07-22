package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// str, _ := ReadTxtToString("SmartFind_0A33908.mhtml")
	// // fmt.Println(strings.Count(str, "\n"))
	// str = strings.ReplaceAll(str, "=\r\n", "")
	// fmt.Println(str)
	// fmt.Println(strings.Count(str, "=\r"))
	// fmt.Println(strings.Count(str, "=\n"))
	// fmt.Println(strings.Count(str, "=\r\n"))
	// fmt.Println(strings.Count(str, "=\r\n\r"))
	// fmt.Println(strings.Count(str, "=\r\n\n"))
	// // fmt.Println(strings.Count(str, "=\r\r"))
	// return

	LowProfileModelList := PrintModes("SmartFind_0A33908.mht")
	SuperLowProfileModelList := PrintModes("SmartFind_4XH0L55146.mht")
	ThinkPad3mmModelList := PrintModes("SmartFind_4XH0X88960.mht")

	fmt.Printf("LowProfileModelList      len = %d\n", len(LowProfileModelList))
	fmt.Printf("SuperLowProfileModelList len = %d\n", len(SuperLowProfileModelList))
	fmt.Printf("ThinkPad3mmModelList len = %d\n", len(ThinkPad3mmModelList))

	ModelListToMarkdown(LowProfileModelList, "CompatibleList-LowProfile.md")
	ModelListToMarkdown(SuperLowProfileModelList, "CompatibleList-SuperLowProfile.md")
	ModelListToMarkdown(ThinkPad3mmModelList, "CompatibleList-ThinkPad3mm.md")
}

func ModelListToMarkdown(modelList []string, fileName string) {
	f, _ := os.Create("../" + fileName)
	defer f.Close()

	f.WriteString(strings.Join(modelList, "\n"))
}

func PrintModes(filePath string) []string {
	var modelList []string
	fileContent, _ := ReadTxtToString(filePath)

	fileContent = strings.ReplaceAll(fileContent, "=\r\n", "")

	//只取 tbody 區塊
	fisrtTbodyIdx := strings.Index(fileContent, "<tbody>")
	lastTbodyIdx := strings.LastIndex(fileContent, "</tbody>")
	fileContent = fileContent[fisrtTbodyIdx:lastTbodyIdx]

	//依照關鍵字開頭分割
	splitedFileContent := strings.Split(fileContent, "<td rowspan=3D\"1\">")

	for _, txtline := range splitedFileContent {
		cutIdx := strings.IndexFunc(txtline, func(r rune) bool { return '<' == r })
		if -1 != cutIdx {
			txtline = txtline[:cutIdx]
		}

		if 0 == len(txtline) {
			continue
		}

		fmt.Printf("%s\n", txtline)
		modelList = append(modelList, fmt.Sprintf("- %s", txtline))
	}

	return modelList
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
