package main

import (
	"log"

	environment "github.com/Wallify/env"
	"github.com/Wallify/sniffing"
)

func main() {
	env := environment.NewDefault()

	builder, err := sniffing.NewBuilder(env)
	if err != nil {
		log.Fatal("could not create builder error: %s", err)
	}

	builder.Build().Run()
}
