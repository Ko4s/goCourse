package topic2
import (
	"bufio"
	"log"
	"os"
)

func OpenFile(filePath string) string {

	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		log.Fatalln(err)
	}
	result := make([]byte, 90)
	file.Read(result)

	return string(result)
}

func ReadLinesFromFile(filePath string) []string {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)

	results := []string{}
	for scanner.Scan() {
		results = append(results, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return results
}
