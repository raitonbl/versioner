package main

import (
	"fmt"
	"github.com/raitonbl/versioner/cmd/options"
	"github.com/raitonbl/versioner/internal"
	"github.com/thatisuday/commando"
	"github.com/thoas/go-funk"
)

func main() {
	builder := internal.Builder{}

	commando.SetExecutableName("versioner").SetVersion("2.0.0").
		SetDescription("allows to read or modify version on source code")

	keys := funk.Keys(builder.GetReaders())

	commando.Register("get").SetShortDescription("reads version from version file").
		SetDescription("This command reads a version file and displays the version").
		AddFlag("type", fmt.Sprintf("indicates the type of version file to be read.\nOptions: %o", keys), commando.String, "maven").
		AddArgument("file", "indicates the version file path", "").
		SetAction(options.ReadVersion(builder.GetReaders()))

	keys = funk.Keys(builder.GetEditors())

	commando.Register("set").SetShortDescription("updates the version for version file").
		SetDescription("This command updates the version file on a specific version file").
		AddFlag("type", fmt.Sprintf("indicates the type of version file to be written.\nOptions: %o", keys), commando.String, "maven").
		AddFlag("version", "indicates the version that should override the version file version.", commando.String, "1.0.0").
		AddArgument("file", "indicates the version file path", "").
		SetAction(options.EditVersion(builder.GetEditors()))

	keys = funk.Keys(builder.GetEnvironments())

	commando.Register("stamp").SetShortDescription("generates the stamp for the specific environment").
		SetDescription("This command generates the a stamp for the specific environment").
		AddFlag("environment", fmt.Sprintf("indicates the environment the stamp is to be generated.\noptions: %o", keys), commando.String, "github").
		SetAction(options.MakeStamp(builder.GetEnvironments()))

	// parse command-line arguments
	commando.Parse(nil)
}
