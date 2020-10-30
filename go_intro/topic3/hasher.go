package topic3

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

type hashedString string

//Hasher struct
type Hasher struct {
	inputFile   string
	outputFile  string
	lines       []string
	hashedLines []hashedString
}

// go mod ini nazwe_modulu -> github/cos
// main.go
//folder -> jakas_nazwa -> hasher.go -> package jakas_nazwe

//NewHasher constructor
func NewHasher(in, out string) *Hasher {
	return &Hasher{
		inputFile:  in,
		outputFile: out,
	}
}

//ReadLines read lines from in file
func (h *Hasher) ReadLines() error {
	file, err := os.Open(h.inputFile)
	defer file.Close()

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

	results := []string{}
	for scanner.Scan() {
		results = append(results, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	h.lines = results //do mojego strcuta Hasher zapisałem sobie results

	return nil
}

func (h *Hasher) HashLines() error {

	for _, line := range h.lines {
		hash, err := h.hashLine(line)

		if err != nil {
			return err
		}
		h.hashedLines = append(h.hashedLines, hash)
	}
	return nil
}

func (h *Hasher) SaveToFile() error {

	out, err := os.OpenFile(h.outputFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	defer out.Close()

	if err != nil {
		return err
	}

	for _, line := range h.hashedLines {
		_, err := out.Write([]byte(line + "\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

//private methods
func (h *Hasher) hashLine(line string) (hashedString, error) {
	hash := sha256.New()

	_, err := hash.Write([]byte(line))

	if err != nil {
		return "", err
	}

	hashedBytes := hash.Sum(nil)

	return hashedString(hex.EncodeToString(hashedBytes)), nil
}
