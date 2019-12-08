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
const EXPECTED_RESULT = 19690720
const MAX_NUMBER = 99

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

		for noun := 0; noun <= MAX_NUMBER; noun++ {
			for verb := 0; verb <= MAX_NUMBER; verb++ {
				result := evaluate(noun, verb, input)
				if result[0] == EXPECTED_RESULT {
					fmt.Println("Combination Found!")
					fmt.Println("Noun: ", noun)
					fmt.Println("Verb: ", verb)
					break
				}
			}
		}
	}
}

func evaluate(noun int, verb int, originalInput []int) []int {
	input := make([]int, len(originalInput))
	copy(input, originalInput)
	input[1] = noun
	input[2] = verb
	result := executeProgram(input)
	return result
}

func executeProgram(input []int) []int {
	pc := 0
	for pc < len(input) {
		opcode := input[pc]
		if HALT == opcode {
			break
		}
		readPos1 := pc + 1
		read1 := input[readPos1]
		readPos2 := pc + 2
		read2 := input[readPos2]
		storePos := pc + 3
		store := input[storePos]
		if ADD == opcode {
			input[store] = input[read1] + input[read2]
		} else if MULTIPLY == opcode {
			input[store] = input[read1] * input[read2]
		} else {
			fmt.Println("Unknown opcode")
			break
		}
		pc += 4
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
