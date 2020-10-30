package topic2

//HashFile return hased string of certain file
func HashFile(filePath string) ([]string, error) {
	//Uwaga to nie jest wszystko
	lines := ReadFileByLine(filePath)
	
	return lines, nil
}
