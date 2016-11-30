package main

import (
	"log"
	"plug1"
)

func main() {
	plug, err := plug1.Open("testplugin.so")
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	plugFunc, err := plug.Lookup("Run")
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	plugFunc.(func())()
}
