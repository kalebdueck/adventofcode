package helpers

import (
	"bufio"
	"log"
	"os"
)

func FetchInputs(fileName string) []string {
	var inputs []string
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, value)
	}

	return inputs
}
