package common

import (
	"github.com/raitonbl/versioner/internal/environment"
	"github.com/raitonbl/versioner/internal/manager"
	"github.com/raitonbl/versioner/pkg"
)

var managers = map[string]pkg.Manager{
	"maven": manager.Maven{},
}
var environments = map[string]pkg.GitEnvironment{
	"github": environment.Github{},
}

func GetManagers() map[string]pkg.Manager {
	return managers
}

func GetEnvironments() map[string]pkg.GitEnvironment {
	return environments
}
