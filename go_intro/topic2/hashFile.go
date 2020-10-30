package topic2

import (
	"log"
	"os"
)

func HashFile(pathFile string) ([]string, error) {
	lines := ReadLinesFromFile(pathFile)

	res := []string{}
	for _, line := range lines {
		hash, err := HashString(line)

		if err != nil {
			return res, err
		}

		res = append(res, hash)
	}

	return res, nil
}

func WriteToFile(readerFile, outputFile string) {
	hashedLines, err := HashFile(readerFile)

	if err != nil {
		log.Fatalln(err)
	}

	out, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer out.Close()

	if err != nil {
		log.Fatalln(err)
	}

	for _, line := range hashedLines {
		_, err := out.WriteString(line + "\n")
		if err != nil {
			log.Fatalln(err)
		}
	}
}
