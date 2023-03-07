package options

import (
	"errors"
	"github.com/raitonbl/versioner/internal/common"
	"github.com/raitonbl/versioner/pkg"
	"github.com/thatisuday/commando"
	"github.com/thoas/go-funk"
	"os"
)

func SetVersion(s map[string]pkg.PackageManager, isStampAware bool, opts map[string]interface{}) func(map[string]commando.ArgValue, map[string]commando.FlagValue) {
	return func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {

		file := args["file"].Value
		object := args["object"].Value
		runtime := args["runtime"].Value
		value, problem := flags["version"].GetString()

		if problem != nil {
			common.DoExit(problem)
		}

		m := s[runtime]

		if m == nil {
			common.DoExit(errors.New("runtime[name='" + runtime + "'] isn't supported"))
		}

		if !funk.Contains(m.GetSupportTypes(), object) {
			common.DoExit(errors.New("runtime[name='" + runtime + "'] doesn't support object[name='" + object + "']"))
		}

		if _, err := os.Stat(file); os.IsNotExist(err) {
			common.DoExit(errors.New(file + " doesn't exist or cannot be opened"))
		}

		v := value

		if isStampAware {
			stamp, err := getStamp(opts["environments"].(map[string]pkg.GitEnvironment), flags)

			if err != nil {
				common.DoExit(err)
			}

			v = v + "-" + stamp
		}

		err := m.SetVersion(object, file, v)

		if err != nil {
			common.DoExit(err)
		}

		os.Exit(0)
	}
}
