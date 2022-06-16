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
			common.Exit(err)
		}

		fmt.Println(v)
		os.Exit(0)
	}
}

func getStamp(cache map[string]pkg.GitEnvironment, flags map[string]commando.FlagValue) (string, error) {
	environment, _ := flags["environment"].GetString()

	var env pkg.GitEnvironment = nil

	for _, current := range cache {
		if current.IsSupported(environment) {
			env = current
			break
		}
	}

	if env == nil {
		return "", errors.New("environment[" + environment + "] isn't supported")
	}

	branchName := ""
	stamp := "SNAPSHOT"

	isPush, problem := env.IsTriggeredByPush()

	if problem != nil {
		return "", problem
	}

	isPullRequest, problem := env.IsTriggeredByPullRequest()

	if problem != nil {
		return "", problem
	}

	if isPush || isPullRequest {
		b, prob := env.GetTargetBranch()

		if prob != nil {
			return "", prob
		}

		if isPush && (strings.HasPrefix(branchName, "release/") || strings.HasPrefix(branchName, "hotfix/")) {
			stamp = "PRERELEASE"
		}

		branchName = b
	}

	if isPush && branchName == env.GetDefaultBranch() {
		return "RELEASE", nil
	}

	pipelineId, prob := env.GetPipeline()

	if prob != nil {
		return "", prob
	}

	return pipelineId + "-" + stamp, nil
}
