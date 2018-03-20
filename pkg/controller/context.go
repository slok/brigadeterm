package controller

import (
	"time"
)

type Project struct {
	Name      string
	ID        string
	LastBuild *Build
}

// ProjectListPageContext has the required information to
// render a project list page.
type ProjectListPageContext struct {
	Projects []*Project
	Error    error
}

type Build struct {
	ID         string
	Version    string
	Running    bool
	FinishedOK bool
	EventType  string
	Started    time.Time
	Ended      time.Time
}

// ProjectBuildListPageContext has the required information to
// render a project build list page.
type ProjectBuildListPageContext struct {
	ProjectName string
	ProjectURL  string
	ProjectNS   string

	Builds []*Build
	Error  error
}

type Job struct {
	ID         string
	Name       string
	Image      string
	Running    bool
	FinishedOK bool
	Started    time.Time
	Ended      time.Time
}

// BuildJobListPageContext has the required information to
// render a build job list page.
type BuildJobListPageContext struct {
	BuildInfo *Build
	Jobs      []*Job
	Error     error
}

// JobLogPageContext has the required information to
// render a job log page.
type JobLogPageContext struct {
	Job   *Job
	Log   []byte
	Error error
}
