package main

import (
	"fmt"
	"github.com/thatisuday/commando"
	"os"
)

type Inspector interface {
	IsSupported(value string) bool
	ReadVersion(value string) (string, error)
	GetVersionFile() string
}

func ReadVersion(array []Inspector) func(map[string]commando.ArgValue, map[string]commando.FlagValue) {
	return func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {

		pathToFile := args["file"].Value
		fileType, _ := flags["type"].GetString()

		if _, err := os.Stat(pathToFile); os.IsNotExist(err) {
			fmt.Println(pathToFile + " doesn't exist or cannot be opened")
			os.Exit(400)
		}

		var inspector Inspector = nil

		for index := 0; index < len(array); index++ {
			current := array[index]

			if current.IsSupported(fileType) {
				inspector = current
			}

		}

		if inspector == nil {
			fmt.Println(fileType + " isn't supported")
			os.Exit(401)
		}

		version, err := inspector.ReadVersion(pathToFile)

		if err != nil {
			fmt.Println(err)
			os.Exit(402)
		}

		fmt.Println(version)
		os.Exit(0)
	}
}
