package options

import (
	"fmt"
	"github.com/raitonbl/versioner/pkg"
	"github.com/thatisuday/commando"
	"os"
)

func ReadVersion(cache map[string]pkg.Reader) func(map[string]commando.ArgValue, map[string]commando.FlagValue) {
	return func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {

		pathToFile := args["file"].Value
		fileType, _ := flags["type"].GetString()

		if _, err := os.Stat(pathToFile); os.IsNotExist(err) {
			fmt.Println(pathToFile + " doesn't exist or cannot be opened")
			os.Exit(400)
		}

		var inspector pkg.Reader = nil

		for _, current := range cache {
			if current.IsSupported(fileType) {
				inspector = current
				break
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
