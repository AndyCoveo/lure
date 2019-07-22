package provider

type PullRequest interface {
	GetSourceBranchName() string
	GetTitle() string
	GetState() string
	GetId() int
}
