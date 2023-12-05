package Common

import (
	"bufio"
	"os"
)

func ReadLines(filePath string) (error, []string) {
	r := make([]string, 0)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, r
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err, r
	}

	return nil, r
}
