package environments

import (
	"errors"
	"os"
	"strings"
)

type Github struct {
}

func (g Github) GetPipeline() (string, error) {

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

func (g Github) GetTargetBranch() (string, error) {

	isValid, err := g.IsTriggeredByPush()

	if err != nil {
		return "", err
	}

	if !isValid {

		isValid, err = g.IsTriggeredByPullRequest()

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

func (g Github) IsSupported(value string) bool {
	return value == "github"
}

func (g Github) IsTriggeredByPush() (bool, error) {
	githubRef := os.Getenv("GITHUB_REF")

	if githubRef == "" {
		return false, errors.New("GITHUB_REF not available")
	}

	return strings.HasPrefix(githubRef, "refs/heads/"), nil
}

func (g Github) IsTriggeredByPullRequest() (bool, error) {
	githubRef := os.Getenv("GITHUB_REF")

	if githubRef == "" {
		return false, errors.New("GITHUB_REF not available")
	}

	return strings.HasPrefix(githubRef, "refs/pull/"), nil
}

func (g Github) GetSourceBranch() (string, error) {
	isPR, err := g.IsTriggeredByPullRequest()

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

func (g Github) IsTriggeredByTag() (bool, error) {
	githubRef := os.Getenv("GITHUB_REF")

	if githubRef == "" {
		return false, errors.New("GITHUB_REF not available")
	}

	return strings.HasPrefix(githubRef, "refs/tags/"), nil
}
