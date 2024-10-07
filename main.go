package main

import (
	"firstgo/firstpkg"
	"fmt"
	"os"

	"flag"

	"github.com/gobuffalo/envy"
)

func main() {

	env := flag.String("env", "prod", "Environment")
	flag.Parse()
	env_prod := ".env." + *env
	envy.Load(env_prod)
	bearer := "Bearer " + os.Getenv("TOKEN")
	fmt.Println("Token in main: ", bearer)
	firstpkg.Called(bearer, os.Getenv("DEL"), os.Getenv("MOD"))
}
