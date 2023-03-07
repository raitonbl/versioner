package options

import (
	"errors"
	"fmt"
	"github.com/raitonbl/versioner/internal/common"
	"github.com/raitonbl/versioner/pkg"
	"github.com/thatisuday/commando"
	"os"
	"strings"
)

func MakeStamp(cache map[string]pkg.GitEnvironment) func(map[string]commando.ArgValue, map[string]commando.FlagValue) {
	return func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {

		v, err := getStamp(cache, flags)

		if err != nil {
			common.DoExit(err)
		}

		fmt.Println(v)
		os.Exit(0)
	}
}

func getStamp(cache map[string]pkg.GitEnvironment, flags map[string]commando.FlagValue) (string, error) {
	environment, _ := flags["environment"].GetString()

	var targetEnvironment pkg.GitEnvironment = nil

	for _, current := range cache {
		if current.GetType() == environment {
			targetEnvironment = current
			break
		}
	}

	if targetEnvironment == nil {
		return "", errors.New("environment[" + environment + "] isn't supported")
	}

	stamp := ""
	branchName, problem := targetEnvironment.GetBranch()
	if problem != nil {
		return "", problem
	}
	isPush, problem := targetEnvironment.IsPushEvent()
	if problem != nil {
		return "", problem
	}
	if isPush {
		if strings.HasPrefix(branchName, "release/") || strings.HasPrefix(branchName, "hotfix/") {
			stamp = "PRERELEASE"
		} else if branchName == targetEnvironment.GetDefaultBranch() {
			return "RELEASE", nil
		}
	} else {
		stamp = "SNAPSHOT"
	}

	pipelineId, prob := targetEnvironment.GetPipelineId()

	if prob != nil {
		return "", prob
	}

	return fmt.Sprintf("%s+%s", stamp, pipelineId), nil
}
