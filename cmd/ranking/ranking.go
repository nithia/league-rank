package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"league-rank/pkg/league"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	lines, err := readInput(in)
	if err != nil {
		log.Fatal(err)
	}

	table := league.RankTable{}
	table.UpdateAll(lines)

	fmt.Print(table)
}

func readInput(in *bufio.Reader) ([]string, error) {
	lines := make([]string, 0, 10)

	for {
		line, err := in.ReadString('\n')
		if err != nil {
			// io.EOF is expected, anything else
			// should be reported
			if err != io.EOF {
				return nil, err
			}
			if line != "" {
				lines = append(lines, line)
			}
			break
		}
		lines = append(lines, line)
	}

	return lines, nil
}
