package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	argsWithProg := os.Args
	filename := argsWithProg[1]
	if len(filename) > 0 {
		f, err := os.Open(filename)
		if err != nil {
			log.Fatalf("Problem opening file : %s", err)
		}

		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)
		var lines []string

		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		f.Close()
		for _, v := range lines {
			fmt.Println(v)
		}
	}

}
