package main

import (
	"fmt"

	"firstgo/firstpkg"

	"github.com/gobuffalo/envy"
)

func main() {

	fmt.Println("Token in main: ", envy.Get("TOKEN", "default"))

	firstpkg.Called((envy.Get("TOKEN", "default")))
}
