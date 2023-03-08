package main

import (
	_ "embed"
	"fmt"
	"github.com/raitonbl/versioner/cmd/options"
	"github.com/raitonbl/versioner/internal"
	"github.com/raitonbl/versioner/pkg"
	"github.com/thatisuday/commando"
	"github.com/thoas/go-funk"
)

func main() {
	packageManagers := internal.GetManagers()
	gitEnvironments := internal.GetEnvironments()
	gitOpts := funk.Keys(gitEnvironments)
	typeOpts, managerOpts := getManagerOptions(packageManagers)
	commando.SetExecutableName("versioner").SetVersion("3.0.1").
		SetDescription("allows to read or modify version on source code")
	commando.Register("get").
		SetShortDescription("reads version from version file").
		SetDescription("This command reads a version from a file and displays it").
		AddArgument("runtime", fmt.Sprintf("indicates the type of object that will be retrieved.\noptions: %s", managerOpts), "").
		AddArgument("object", fmt.Sprintf("indicates the type of object that will be retrieved.\noptions: %s", typeOpts), "").
		AddArgument("file", "indicates the version file path", "").
		SetAction(options.GetVersion(packageManagers, map[string]interface{}{
			"types":    typeOpts,
			"managers": managerOpts,
		}))
	commando.Register("set").
		SetShortDescription("updates the version for version file").
		SetDescription("This command updates the version file on a specific version file").
		AddFlag("version", "indicates the version that should override the version file version.", commando.String, "1.0.0").
		AddArgument("runtime", fmt.Sprintf("indicates the type of object that will be updated.\noptions: %s", managerOpts), "").
		AddArgument("object", fmt.Sprintf("indicates the type of object that will be updated.\noptions: %s", typeOpts), "").
		AddArgument("file", "indicates the version file path", "").
		SetAction(options.SetVersion(packageManagers, false, map[string]interface{}{
			"types":        typeOpts,
			"managers":     managerOpts,
			"environments": gitEnvironments,
		}))
	commando.Register("stamp").SetShortDescription("generates the stamp for the specific environment").
		SetDescription("This command generates the a stamp for the specific environment").
		AddFlag("environment", fmt.Sprintf("indicates the environment the stamp is to be generated.\noptions: %o", gitOpts), commando.String, "github").
		SetAction(options.MakeStamp(gitEnvironments))
	commando.Register("set-stamped-version").
		SetShortDescription("updates the version for the specified and attach's the stamp").
		SetDescription("updates the version for the specified and attach's the stamp").
		AddFlag("version", "indicates the version that should override the version file version.", commando.String, "1.0.0").
		AddFlag("environment", fmt.Sprintf("indicates the environment the stamp is to be generated.\noptions: %o", gitOpts), commando.String, "github").
		AddArgument("runtime", fmt.Sprintf("indicates the type of object that will be updated.\noptions: %s", managerOpts), "").
		AddArgument("object", fmt.Sprintf("indicates the type of object that will be updated.\noptions: %s", typeOpts), "").
		AddArgument("file", "indicates the version file path", "").
		SetAction(options.SetVersion(packageManagers, true, map[string]interface{}{
			"types":        typeOpts,
			"managers":     managerOpts,
			"environments": gitEnvironments,
		}))
	commando.Register("get-stamped-version").
		SetShortDescription("updates the version for the specified and attach's the stamp").
		SetDescription("updates the version for the specified and attach's the stamp").
		AddFlag("environment", fmt.Sprintf("indicates the environment the stamp is to be generated.\noptions: %o", gitOpts), commando.String, "github").
		AddArgument("runtime", fmt.Sprintf("indicates the type of object that will be updated.\noptions: %s", managerOpts), "").
		AddArgument("object", fmt.Sprintf("indicates the type of object that will be updated.\noptions: %s", typeOpts), "").
		AddArgument("file", "indicates the version file path", "").
		SetAction(options.GetStampedVersion(gitEnvironments, packageManagers, map[string]interface{}{
			"types":        typeOpts,
			"managers":     managerOpts,
			"environments": gitEnvironments,
		}))
	commando.Parse(nil)
}

func getManagerOptions(managers map[string]pkg.PackageManager) ([]string, interface{}) {
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
