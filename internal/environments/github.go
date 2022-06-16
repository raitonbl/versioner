package environments

import (
	"errors"
	"os"
	"strings"
)

type Github struct {
}

func (instance Github) GetDefaultBranch() string {
	return "main"
}

func (instance Github) GetPipeline() (string, error) {

	githubRunId := os.Getenv("GITHUB_RUN_ID")

	if githubRunId == "" {
		return "", errors.New("GITHUB_RUN_ID not available")
	}

	githubRunNumber := os.Getenv("GITHUB_RUN_NUMBER")

	if githubRunNumber == "" {
		return "", errors.New("GITHUB_RUN_NUMBER not available")
	}

	return githubRunId + "." + githubRunNumber, nil
}

func (instance Github) IsSupported(value string) bool {
	return value == "github"
}

func (instance Github) GetSourceBranch() (string, error) {
	isPR, err := instance.IsTriggeredByPullRequest()

	if err != nil {
		return "", err
	}

	if !isPR {
		return "", errors.New("GITHUB_HEAD_REF not available")
	}

	githubHeadRef := os.Getenv("GITHUB_HEAD_REF")

	if githubHeadRef == "" {
		return "", errors.New("GITHUB_HEAD_REF not available")
	}

	return githubHeadRef, nil

}

func (instance Github) IsTriggeredByPush() (bool, error) {
	githubRef := os.Getenv("GITHUB_REF")

	if githubRef == "" {
		return false, errors.New("GITHUB_REF not available")
	}

	return strings.HasPrefix(githubRef, "refs/heads/"), nil
}

func (instance Github) GetTargetBranch() (string, error) {

	isValid, err := instance.IsTriggeredByPush()

	if err != nil {
		return "", err
	}

	if !isValid {

		isValid, err = instance.IsTriggeredByPullRequest()

		if err != nil {
			return "", err
		}

	}

	if !isValid {
		return "", errors.New("GIT_REF not supported")
	}

	githubRefName := os.Getenv("GITHUB_REF_NAME")

	if githubRefName == "" {
		return "", errors.New("GITHUB_REF_NAME not available")
	}

	return githubRefName, nil
}

func (instance Github) IsTriggeredByPullRequest() (bool, error) {
	githubRef := os.Getenv("GITHUB_REF")

	if githubRef == "" {
		return false, errors.New("GITHUB_REF not available")
	}

	return strings.HasPrefix(githubRef, "refs/pull/"), nil
}
