package firstpkg

import (
	"fmt"
)

func Called(toke string, delidx string, modidx string) {
	tok := "this-" + toke + "-will be used to delete " + delidx + " and modify " + modidx
	fmt.Printf("Passed from main: %s", tok)
}
