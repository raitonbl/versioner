package pkg

type GitEnvironment interface {
	GetPipeline() (string, error)
	IsSupported(value string) bool
	GetSourceBranch() (string, error)
	GetTargetBranch() (string, error)

	IsTriggeredByTag() (bool, error)
	IsTriggeredByPush() (bool, error)
	IsTriggeredByPullRequest() (bool, error)
}
