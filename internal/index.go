package internal

import (
	"github.com/raitonbl/versioner/internal/environment"
	"github.com/raitonbl/versioner/internal/manager"
	"github.com/raitonbl/versioner/pkg"
)

var managers = map[string]pkg.PackageManager{
	"oas3":   manager.OAS3{},
	"helm":   manager.Helm{},
	"maven":  manager.Maven{},
	"nodejs": manager.Nodejs{},
}
var environments = map[string]pkg.GitEnvironment{
	"github": environment.Github{},
}

func GetManagers() map[string]pkg.PackageManager {
	return managers
}

func GetEnvironments() map[string]pkg.GitEnvironment {
	return environments
}
