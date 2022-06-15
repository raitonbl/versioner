package options

import (
	"errors"
	"fmt"
	"github.com/raitonbl/versioner/internal/commons"
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
			commons.Exit(errors.New("environment[" + environment + "] isn't supported"))
		}

		branchName := ""
		stamp := "SNAPSHOT"
		addPipelineId := true

		isTag, err := env.IsTriggeredByTag()

		if err != nil {
			commons.Exit(err)
		}

		if isTag {
			stamp = "RELEASE"
			addPipelineId = false
		} else {

			isPush, pr := env.IsTriggeredByPush()

			if pr != nil {
				commons.Exit(pr)
			}

			isPullRequest, pr := env.IsTriggeredByPullRequest()

			if pr != nil {
				commons.Exit(pr)
			}

			if isPush || isPullRequest {
				b, prob := env.GetTargetBranch()

				if prob != nil {
					commons.Exit(prob)
				}

				if isPush && (strings.HasPrefix(branchName, "release/") || strings.HasPrefix(branchName, "hotfix/")) {
					stamp = "PRERELEASE"
				}

				branchName = b
			}
		}

		if addPipelineId {
			pipelineId, prob := env.GetPipeline()

			if prob != nil {
				commons.Exit(prob)
			}

			fmt.Println(pipelineId + "-" + stamp)
		} else {
			fmt.Println(stamp)
		}

		os.Exit(0)
	}
}
