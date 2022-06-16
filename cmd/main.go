package main

import (
	_ "embed"
	"fmt"
	"github.com/raitonbl/versioner/cmd/options"
	"github.com/raitonbl/versioner/internal/common"
	"github.com/raitonbl/versioner/pkg"
	"github.com/thatisuday/commando"
	"github.com/thoas/go-funk"
)

//go:embed version.txt
var version string

func main() {

	managers := common.GetManagers()
	gitEnvironments := common.GetEnvironments()

	gitOpts := funk.Keys(gitEnvironments)
	typeOpts, managerOpts := getManagerOptions(managers)

	commando.SetExecutableName("versioner").SetVersion(version).
		SetDescription("allows to read or modify version on source code")

	commando.Register("get").
		SetShortDescription("reads version from version file").
		SetDescription("This command reads a version from a file and displays it").
		AddArgument("runtime", fmt.Sprintf("indicates the type of object that will be retrieved.\noptions: %s", managerOpts), "").
		AddArgument("object", fmt.Sprintf("indicates the type of object that will be retrieved.\noptions: %s", typeOpts), "").
		AddArgument("file", "indicates the version file path", "").
		SetAction(options.GetVersion(managers, map[string]interface{}{
			"types":    typeOpts,
			"managers": managerOpts,
		}))

	commando.Register("set").
		SetShortDescription("updates the version for version file").
		SetDescription("This command updates the version file on a specific version file").
		AddFlag("value", "indicates the version that should override the version file version.", commando.String, "1.0.0").
		AddArgument("runtime", fmt.Sprintf("indicates the type of object that will be updated.\noptions: %s", managerOpts), "").
		AddArgument("object", fmt.Sprintf("indicates the type of object that will be updated.\noptions: %s", typeOpts), "").
		AddArgument("file", "indicates the version file path", "").
		SetAction(options.SetVersion(managers, map[string]interface{}{
			"types":    typeOpts,
			"managers": managerOpts,
		}))

	commando.Register("stamp").SetShortDescription("generates the stamp for the specific environment").
		SetDescription("This command generates the a stamp for the specific environment").
		AddFlag("environment", fmt.Sprintf("indicates the environment the stamp is to be generated.\noptions: %o", gitOpts), commando.String, "github").
		SetAction(options.MakeStamp(gitEnvironments))

	// parse command-line arguments
	commando.Parse(nil)
}

func getManagerOptions(managers map[string]pkg.Manager) ([]string, interface{}) {
	typeOpts := make([]string, 0)
	managerOpts := funk.Keys(managers)

	for _, manager := range managers {
		for _, each := range manager.GetSupportTypes() {
			if !funk.Contains(typeOpts, each) {
				typeOpts = append(typeOpts, each)
			}
		}
	}

	return typeOpts, managerOpts
}
