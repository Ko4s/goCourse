package topic2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


func OpenFile(filePath string) {

	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		log.Fatalf("%v", err)
	}

	fileContent := make([]byte, 8)
	file.Read(fileContent)

	fmt.Println(string(fileContent))
}

func ReadFileByLine(filePath string) []string {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		log.Fatalf("%v", err)
	}

	scanner := bufio.NewScanner(file)
	result := []string{}
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("%v", err)
	}

	return result
}
