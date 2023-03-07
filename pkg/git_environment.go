package pkg

type GitEnvironment interface {
	GetDefaultBranch() string
	GetPipelineId() (string, error)
	GetBranch() (string, error)
	GetType() string
	IsPushEvent() (bool, error)
	IsPullRequestEvent() (bool, error)
}
