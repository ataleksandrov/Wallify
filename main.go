package main

import (
	"fmt"

	environment "github.com/Wallify/env"
)

func main() {
	env := environment.NewDefault()
	fmt.Println(*env.HC)
	fmt.Println(*env.RC)
}
