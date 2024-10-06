package firstpkg

import (
	"fmt"
)

func Called(toke string) {
	tok := "this-" + toke + "-came from main"
	fmt.Printf("Passed from main: %s", tok)
}
