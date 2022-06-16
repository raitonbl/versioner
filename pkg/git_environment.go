package pkg

type GitEnvironment interface {
	GetDefaultBranch() string

	GetPipeline() (string, error)
	IsSupported(value string) bool
	GetSourceBranch() (string, error)
	GetTargetBranch() (string, error)

	IsTriggeredByPush() (bool, error)
	IsTriggeredByPullRequest() (bool, error)
}
