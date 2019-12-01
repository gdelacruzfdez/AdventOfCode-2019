package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	totalFuel := 0

	for scanner.Scan() {
		weight, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}

		totalFuel += getModuleFuel(weight)
	}

	fmt.Println("Total Fuel for modules: ", totalFuel)
}

func getModuleFuel(weight int) int {
	totalModuleFuel := 0
	for weight > 0 {
		fuel := calculateFuelFromWeight(weight)
		if fuel > 0 {
			totalModuleFuel += fuel
		}
		weight = fuel
	}
	return totalModuleFuel
}

func calculateFuelFromWeight(weight int) int {
	return int(weight/3) - 2
}
