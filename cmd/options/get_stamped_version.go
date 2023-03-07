package options

import (
	"errors"
	"fmt"
	"github.com/raitonbl/versioner/internal/common"
	"github.com/raitonbl/versioner/pkg"
	"github.com/thatisuday/commando"
	"github.com/thoas/go-funk"
	"os"
)

func GetStampedVersion(gitEnvironments map[string]pkg.GitEnvironment, s map[string]pkg.PackageManager, _ map[string]interface{}) func(map[string]commando.ArgValue, map[string]commando.FlagValue) {
	return func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
		file := args["file"].Value
		object := args["object"].Value
		runtime := args["runtime"].Value
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
		version, err := m.GetVersion(object, file)
		if err != nil {
			common.DoExit(err)
		}
		stamp, err := getStamp(gitEnvironments, flags)
		if err != nil {
			common.DoExit(err)
		}
		version = fmt.Sprintf("%s-%s", version, stamp)
		fmt.Println(version)
		os.Exit(0)
	}
}
