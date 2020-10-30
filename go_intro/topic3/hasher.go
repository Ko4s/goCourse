package topic3

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

/*
	in -> plik wejsiowy
	out -> plik wyjscia
	line -> []string wszystkich lini z pliku in
	hashedLines -> []hash(mój własny typ) -> tablica zahashowanych stringow
*/

type hashedString string

//Hasher struct
type Hasher struct {
	in          string
	out         string
	lines       []string
	hashedLines []hashedString
}

//NewHasher - Hasher constructor
func NewHasher(in, out string) *Hasher {
	return &Hasher{ //& zwróci mi pointera
		in:  in,
		out: out,
	}
}

//ReadLines reads file(h.in) lines
func (h *Hasher) ReadLines() error {
	file, err := os.Open(h.in)
	defer file.Close()

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	result := []string{}
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	h.lines = result
	return nil //wszystko było ok wiec nie ma błedu
}

//HashLines hashing h.lines
func (h *Hasher) HashLines() error {

	for _, el := range h.lines {
		hash, err := h.hashLine(el)

		if err != nil {
			return nil
		}

		h.hashedLines = append(h.hashedLines, hash)
	}

	return nil
}

// pomocnicza funckja prywatna dla paczki
func (h *Hasher) hashLine(line string) (hashedString, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(line))

	if err != nil {
		return "", err
	}

	hashBytes := hash.Sum(nil)
	return hashedString(hex.EncodeToString(hashBytes)), nil // converutje string -> hashedString
}

//Save is saving hasheLines to file described in h.out
func (h *Hasher) Save() error {
	f, err := os.OpenFile(h.out, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	if err != nil {
		return err
	}

	for _, el := range h.hashedLines {
		if _, err := f.Write([]byte(el + "\n")); err != nil {
			return err
		}
	}

	return nil
}
