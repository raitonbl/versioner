package options

import (
	"errors"
	"github.com/raitonbl/versioner/internal/common"
	"github.com/raitonbl/versioner/pkg"
	"github.com/thatisuday/commando"
	"github.com/thoas/go-funk"
	"os"
)

func SetVersion(s map[string]pkg.Manager, isStampAware bool, opts map[string]interface{}) func(map[string]commando.ArgValue, map[string]commando.FlagValue) {
	return func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {

		file := args["file"].Value
		object := args["object"].Value
		runtime := args["runtime"].Value
		value, _ := flags["value"].GetString()

		var m pkg.Manager = nil

		for _, current := range s {
			t := s[runtime]

			if t != nil {
				m = current
				break
			}
		}

		if m == nil {
			common.Exit(errors.New("runtime[name='" + runtime + "'] isn't supported"))
		}

		if !funk.Contains(m.GetSupportTypes(), object) {
			common.Exit(errors.New("runtime[name='" + runtime + "'] doesn't support object[name='" + object + "']"))
		}

		if _, err := os.Stat(file); os.IsNotExist(err) {
			common.Exit(errors.New(file + " doesn't exist or cannot be opened"))
		}

		v := value

		if isStampAware {
			stamp, err := getStamp(opts["environments"].(map[string]pkg.GitEnvironment), flags)

			if err != nil {
				common.Exit(err)
			}

			v = v + "-" + stamp
		}

		err := m.SetVersion(object, file, v)

		if err != nil {
			common.Exit(err)
		}

		os.Exit(0)
	}
}
