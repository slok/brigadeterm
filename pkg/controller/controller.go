package controller

import (
	"fmt"

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
}

type controller struct {
	brigade brigade.Service
}

// NewController returns a new controller.
func NewController(brigade brigade.Service) Controller {
	return &controller{
		brigade: brigade,
	}
}

func (c *controller) ProjectListPageContext() *ProjectListPageContext {
	prjs, err := c.brigade.GetAllProjects()
	if err != nil {
		return &ProjectListPageContext{
			Error: fmt.Errorf("there was an error while getting projects from brigade: %s", err),
		}
	}

	ctxPrjs := make([]*project, len(prjs))
	for i, prj := range prjs {
		ctxPrjs[i] = &project{
			ID:   prj.ID,
			Name: prj.Name,
		}
	}

	return &ProjectListPageContext{
		Projects: ctxPrjs,
	}
}

func (c *controller) ProjectBuildListPageContext(projectID string) *ProjectBuildListPageContext {
	return &ProjectBuildListPageContext{
		Error: fmt.Errorf("not implentented"),
	}
}

func (c *controller) BuildJobListPageContext(projectID string) *BuildJobListPageContext {
	return &BuildJobListPageContext{
		Error: fmt.Errorf("not implentented"),
	}
}
