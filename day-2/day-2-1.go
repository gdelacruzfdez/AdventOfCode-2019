package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const HALT = 99
const ADD = 1
const MULTIPLY = 2

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		inputStrs := strings.Split(line, ",")
		input, err := convertStrArrToIntArr(inputStrs)
		if err != nil {
			panic(err)
		}
		input = restoreGravityAssistProgram(input)
		output := executeProgram(input)
		fmt.Println(output)
	}
}

func restoreGravityAssistProgram(input []int) []int {
	input[1] = 12
	input[2] = 2
	return input
}

func executeProgram(input []int) []int {
	for i := 0; i < len(input); i += 4 {
		opcode := input[i]
		if HALT == opcode {
			break
		}
		readPos1 := i + 1
		read1 := input[readPos1]
		readPos2 := i + 2
		read2 := input[readPos2]
		storePos := i + 3
		store := input[storePos]
		if ADD == opcode {
			input[store] = input[read1] + input[read2]
		}
		if MULTIPLY == opcode {
			input[store] = input[read1] * input[read2]
		}
	}
	return input
}

func convertStrArrToIntArr(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
