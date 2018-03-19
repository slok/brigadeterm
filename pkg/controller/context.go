package controller

import (
	"time"
)

type project struct {
	Name               string
	ID                 string
	LastBuildVersion   string
	LastBuildOK        bool
	LastBuildTime      time.Time
	LastBuildEventType string
}

// ProjectListPageContext has the required information to
// render a project list page.
type ProjectListPageContext struct {
	Projects []*project
	Error    error
}

type build struct {
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

	Builds []*build
	Error  error
}

type job struct {
	ID         string
	Name       string
	Running    bool
	FinishedOK bool
	Started    time.Time
	Ended      time.Time
}

// BuildJobListPageContext has the required information to
// render a build job list page.
type BuildJobListPageContext struct {
	BuildInfo *build
	Jobs      []*job
	Error     error
}
