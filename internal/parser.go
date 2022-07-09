package internal

import (
	"bufio"
	"log"
	"os"
	"time"
)

func Load(filePath string) []time.Duration {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	result := make([]time.Duration, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		parsedDuration, _ := time.ParseDuration(text)
		result = append(result, parsedDuration)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
