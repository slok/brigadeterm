package controller

import (
	"fmt"

	azurebrigade "github.com/Azure/brigade/pkg/brigade"

	brigademodel "github.com/slok/brigadeterm/pkg/model/brigade"
	"github.com/slok/brigadeterm/pkg/service/brigade"
)

// Controller knows what to how to handle the different ui views data
// using the required services and having the logic of each part.
type Controller interface {
	// ProjectListPageContext returns the projectListPage context.
	ProjectListPageContext() *ProjectListPageContext
	// ProjectBuildListContext returns the projectBuildListPage context.
	ProjectBuildListPageContext(projectID string) *ProjectBuildListPageContext
	// BuildJobListPageContext returns the BuildJobListPage context.
	BuildJobListPageContext(buildID string) *BuildJobListPageContext
	// JobLogPageContext returns the JobLogPage context.
	JobLogPageContext(jobID string) *JobLogPageContext
}

type controller struct {
	brigade    brigade.Service
	Controller // For the fakes.
}

// NewController returns a new controller.
func NewController(brigade brigade.Service) Controller {
	return &controller{
		brigade:    brigade,
		Controller: NewFakeController(),
	}
}

func (c *controller) ProjectListPageContext() *ProjectListPageContext {
	prjs, err := c.brigade.GetProjects()
	if err != nil {
		return &ProjectListPageContext{
			Error: fmt.Errorf("there was an error while getting projects from brigade: %s", err),
		}
	}

	ctxPrjs := make([]*Project, len(prjs))
	for i, prj := range prjs {
		lastBuild, err := c.brigade.GetProjectLastBuild(prj.ID)
		if err != nil {
			return &ProjectListPageContext{
				Error: fmt.Errorf("there was an error while getting project %s las build from brigade: %s", prj.ID, err),
			}
		}
		ctxPrjs[i] = &Project{
			ID:        prj.ID,
			Name:      prj.Name,
			LastBuild: c.transformBuild(lastBuild),
		}
	}

	return &ProjectListPageContext{
		Projects: ctxPrjs,
	}
}

func (c *controller) ProjectBuildListPageContext(projectID string) *ProjectBuildListPageContext {
	prj, err := c.brigade.GetProject(projectID)
	if err != nil {
		return &ProjectBuildListPageContext{
			Error: fmt.Errorf("there was an error while getting project from brigade: %s", err),
		}
	}

	builds, err := c.brigade.GetProjectBuilds(prj)
	if err != nil {
		return &ProjectBuildListPageContext{
			Error: fmt.Errorf("there was an error while getting builds from brigade: %s", err),
		}
	}

	ctxBuilds := make([]*Build, len(builds))
	for i, b := range builds {
		ctxBuilds[i] = c.transformBuild(b)
	}

	return &ProjectBuildListPageContext{
		ProjectName: prj.Name,
		ProjectNS:   prj.Kubernetes.Namespace,
		ProjectURL:  prj.Repo.CloneURL,
		Builds:      ctxBuilds,
	}
}

// func (c *controller) BuildJobListPageContext(projectID string) *BuildJobListPageContext {
// 	return &BuildJobListPageContext{
// 		Error: fmt.Errorf("not implentented"),
// 	}
// }

// func (c *controller) JobLogPageContext(jobID string) *JobLogPageContext {
// 	return &JobLogPageContext{
// 		Error: fmt.Errorf("not implentented"),
// 	}
// }

func (c *controller) transformBuild(b *brigademodel.Build) *Build {
	isRunning := false
	ok := false
	switch b.Worker.Status {
	case azurebrigade.JobRunning, azurebrigade.JobPending:
		isRunning = true
	case azurebrigade.JobSucceeded:
		ok = true
	}

	return &Build{
		ID:         b.ID,
		Version:    b.Revision.Commit,
		Running:    isRunning,
		FinishedOK: ok,
		EventType:  b.Type,
		Started:    b.Worker.StartTime,
		Ended:      b.Worker.EndTime,
	}
}
