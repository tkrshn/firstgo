package firstpkg

import (
	"fmt"
)

func Called(toke string) {
	fmt.Printf("Passed from main: %s", toke)
}
