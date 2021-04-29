package Files

import (
	"bufio"
	"io"
	"os"
)

func ReadAllLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	result := make([]string, 0, 300)
	br := bufio.NewReader(file)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		result = append(result, string(a))
	}
	return result, nil
}
