// The colours are created and populated here
// Using structs for the colour types
package colours

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Colour struct {
	colour string
}

func (c Colour) String() string {
	return fmt.Sprintf("%v", c.colour)
}

// return array of Colour objects that can be accessed in Controller
func PopulateColours(x int) []Colour {
	colors := make([]Colour, x)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < len(colors); i++ {
		fmt.Printf("Please enter colour number %d:\n", i+1)
		text, _ := reader.ReadString('\n')
		colourText := strings.TrimSpace(text)
		colors[i] = Colour{colourText}
	}
	return colors
}
