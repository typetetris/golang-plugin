package main

import (
	"log"
	"plugin"
)

func main() {
	plug, err := plugin.Open("testplugin.so")
	if err != nil {
		log.Fatalf("%v\n")
	}
	plugFunc, err := plug.Lookup("Run")
	if err != nil {
		log.Fatalf("%v\n")
	}
	plugFunc.(func())()
}
