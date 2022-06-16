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
		environment, _ := flags["environment"].GetString()

		var env pkg.GitEnvironment = nil

		for _, current := range cache {
			if current.IsSupported(environment) {
				env = current
				break
			}
		}

		if env == nil {
			common.Exit(errors.New("environment[" + environment + "] isn't supported"))
		}

		branchName := ""
		stamp := "SNAPSHOT"

		isPush, problem := env.IsTriggeredByPush()

		if problem != nil {
			common.Exit(problem)
		}

		isPullRequest, problem := env.IsTriggeredByPullRequest()

		if problem != nil {
			common.Exit(problem)
		}

		if isPush || isPullRequest {
			b, prob := env.GetTargetBranch()

			if prob != nil {
				common.Exit(prob)
			}

			if isPush && (strings.HasPrefix(branchName, "release/") || strings.HasPrefix(branchName, "hotfix/")) {
				stamp = "PRERELEASE"
			}

			branchName = b
		}

		if isPush && branchName == env.GetDefaultBranch() {
			fmt.Println("RELEASE")
			os.Exit(0)
		}

		pipelineId, prob := env.GetPipeline()

		if prob != nil {
			common.Exit(prob)
		}

		fmt.Println(pipelineId + "-" + stamp)
		os.Exit(0)
	}
}
