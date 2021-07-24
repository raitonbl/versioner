package main

import (
	"github.com/thatisuday/commando"
)

func main() {

	commando.SetExecutableName("versioner").SetVersion("1.0.0").
		SetDescription("allows to read or modify version on source code")

	commando.Register("read").SetShortDescription("reads version from version file").
		SetDescription("This command reads a version file and displays the version").
		AddFlag("type", "indicates the type of version file to be read", commando.String, "maven").
		AddArgument("file", "indicates the version file path", "./pom.xml").
		SetAction(ReadVersion(createInspectorChain()))

	commando.Register("update").SetShortDescription("updates the version for version file").
		SetDescription("This command updates the version file on a specific version file").
		AddFlag("type", "indicates the type of version file to be written", commando.String, "maven").
		AddFlag("version", "indicates the version that should override the version file version", commando.String, "1.0.0").
		AddArgument("file", "indicates the version file path", "pom.xml").
		SetAction(OnWrite(createEditorChain()))

	// parse command-line arguments
	commando.Parse(nil)

}

func createInspectorChain() []Inspector {
	array := make([]Inspector, 1)
	array[0] = Maven{}
	return array
}

func createEditorChain() []Editor {
	array := make([]Editor, 1)
	array[0] = Maven{}
	return array
}
