package internal

import (
	"github.com/raitonbl/versioner/internal/environments"
	"github.com/raitonbl/versioner/internal/managers"
	"github.com/raitonbl/versioner/pkg"
)

type Builder struct {
	readers      map[string]pkg.Reader
	editors      map[string]pkg.Editor
	environments map[string]pkg.GitEnvironment
}

func (instance *Builder) GetReaders() map[string]pkg.Reader {

	if instance.readers == nil || len(instance.readers) > 0 {
		instance.readers = make(map[string]pkg.Reader)
		instance.readers["maven"] = managers.Maven{}
	}

	return instance.readers
}

func (instance *Builder) GetEditors() map[string]pkg.Editor {

	if instance.editors == nil || len(instance.editors) > 0 {
		instance.editors = make(map[string]pkg.Editor)
		instance.editors["maven"] = managers.Maven{}
	}

	return instance.editors
}

func (instance *Builder) GetEnvironments() map[string]pkg.GitEnvironment {

	if instance.environments == nil || len(instance.environments) > 0 {
		instance.environments = make(map[string]pkg.GitEnvironment)
		instance.environments["github"] = environments.Github{}
	}

	return instance.environments
}
