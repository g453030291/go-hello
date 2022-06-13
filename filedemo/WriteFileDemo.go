package filedemo

import (
	"fmt"
	"os"
)

var lines = []string{
	"Go",
	"is",
	"the",
	"best",
	"programming",
	"language",
	"in",
	"the",
	"world",
}

func WriteFile() {
	file, error := os.Create("file.txt")
	if error != nil {
		fmt.Println(error)
	}
	defer file.Close()

	for i, line := range lines {
		fmt.Println(i)
		file.WriteString("第" + string(i) + "行" + line + "\n")
	}
}
