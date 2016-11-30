package main

import "C"

import (
	"log"
	"plugin"
)

//export RunGoPlugin
func RunGoPlugin() {
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

func main() {
	//for buildmode c-archive
}
