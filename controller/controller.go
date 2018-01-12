// The logic of the program is handled here
package controller

import (
	"Lights/colours"
	"fmt"
	"time"
)

func RunController(numLights, numColours int) {
	if numLights <= 0 {
		fmt.Println("###########################################################################")
		fmt.Println("############################# WARNING!!! ##################################")
		fmt.Println("###########################################################################")
		fmt.Println("### This controller is only designed to work with one or more lights... ###")
		fmt.Println("###########################################################################")
		fmt.Println("###########################################################################")
		fmt.Println("###########################################################################")
	} else {
		// variables used in this part of the controller:
		var isOn bool // initialises to false
		colourObjs := colours.PopulateColours(numColours)
		x := 0

		loc, err := time.LoadLocation("Europe/London")

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for x < len(colourObjs) {
			t := time.Now()
			if isOn {
				isOn = false
			} else {
				fmt.Printf("%s %+v %v\n", t.In(loc).Format("02/01/06 03:04:05 PM Jan"), colourObjs[x], "ON")
				time.Sleep(1 * time.Second)
				isOn = true
				fmt.Printf("%s %+v %v\n", t.In(loc).Format("02/01/06 03:04:05 PM Jan"), colourObjs[x], "OFF")
				time.Sleep(1 * time.Second)
				isOn = false
				x++
			}
		}
	}
}
