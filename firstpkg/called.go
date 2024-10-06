package firstpkg

import (
	"fmt"
)

func Called(tok string) {
	fmt.Println("Passed from main: ", tok)
}
