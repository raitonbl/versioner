package environment

import (
	"errors"
	"os"
	"strings"
)

type Github struct {
}

func (instance Github) GetType() string {
	return "github"
}

func (instance Github) GetDefaultBranch() string {
	return "main"
}

func (instance Github) GetPipelineId() (string, error) {

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

func (instance Github) GetBranch() (string, error) {
	isPR, err := instance.IsPullRequestEvent()

	if err != nil {
		return "", err
	}

	if !isPR {
		return "", errors.New("GITHUB_REF not available")
	}

	githubHeadRef := os.Getenv("GITHUB_REF")

	if githubHeadRef == "" {
		return "", errors.New("GITHUB_REF not available")
	}

	return githubHeadRef, nil

}

func (instance Github) IsPushEvent() (bool, error) {
	githubRef := os.Getenv("GITHUB_REF")

	if githubRef == "" {
		return false, errors.New("GITHUB_REF not available")
	}

	return strings.HasPrefix(githubRef, "refs/heads/"), nil
}

func (instance Github) IsPullRequestEvent() (bool, error) {
	githubRef := os.Getenv("GITHUB_REF")

	if githubRef == "" {
		return false, errors.New("GITHUB_REF not available")
	}

	return strings.HasPrefix(githubRef, "refs/pull/"), nil
}
