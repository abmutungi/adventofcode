package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("measurements.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var measurements []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		measurements = append(measurements, scanner.Text())
	}


	count := 0

	for i := 1; i < len(measurements); i++ {
		if measurements[i-1] < measurements[i] {
			count++
		}		
	}
	fmt.Println(count+1)
}
