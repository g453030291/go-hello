package filedemo

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var path string = "/Users/gemushen/goproject/go-hello/README.md"

func ReadFile() {
	file, error := os.Open(path)

	if error != nil {
		fmt.Println("读取文件报错了:" + error.Error())
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		input, error := reader.ReadString('\n')
		if error == io.EOF {
			return
		}
		fmt.Println(input)
	}
}

func ReadFile2() {
	content, error := os.ReadFile(path)
	if error != nil {
		print(error)
	}
	fmt.Println(string(content))
}
