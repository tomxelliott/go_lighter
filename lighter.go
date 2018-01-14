// The handling of command line arguments are dealt with here
// Once the input has been validated we can call the rest of the program
package main

import (
	"Lights/controller"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var noOfLights int
var noOfColours int

func main() {
	if len(os.Args) > 1 {
		fmt.Println("This program does not take any additional arguments!")
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please enter how many lights you want:")
		valid := false
		for valid == false {
			text, _ := reader.ReadString('\n')
			lightNo := strings.TrimSpace(text)

			if ln, err := strconv.Atoi(lightNo); err != nil {
				fmt.Println(err.Error())
				fmt.Println("WARNING: Please enter a valid number of lights:")
			} else {
				if ln <= 0 {
					fmt.Println("WARNING! Please enter a positive number greater than Zero:")
				} else {
					valid = true
				}
			}

			if valid == true {
				lightNoInt, _ := strconv.Atoi(lightNo)
				noOfLights = lightNoInt
			}
		}

		fmt.Println("Please enter how many colours you want:")

		validColours := false
		for validColours == false {
			text, _ := reader.ReadString('\n')

			colourNo := strings.TrimSpace(text)

			if cn, err := strconv.Atoi(colourNo); err != nil {
				fmt.Println(err.Error())
				fmt.Println("Please enter a valid number of colours:")
			} else if cn <= 0 {
				fmt.Println("Please enter a number greater than ZERO (0) for your number of colours:")
			} else if cn > noOfLights {
				noOfColours = noOfLights
				fmt.Println("Warning!")
				fmt.Println("The number of colours exceeded the number of lights.")
				fmt.Printf("The number of colours has been set to %d.\n", noOfColours)
				validColours = true
			} else {
				coloursNoInt, _ := strconv.Atoi(colourNo)
				noOfColours = coloursNoInt
				validColours = true
			}
		}

		defer fmt.Println("Program quitting...")
		controller.RunController(noOfLights, noOfColours)
		time.Sleep(1 * time.Second)
	}
}
