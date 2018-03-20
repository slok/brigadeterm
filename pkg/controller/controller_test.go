package controller_test

import (
	"testing"
	"time"

	azurebrigade "github.com/Azure/brigade/pkg/brigade"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/slok/brigadeterm/pkg/controller"
	mbrigade "github.com/slok/brigadeterm/pkg/mocks/service/brigade"
	brigademodel "github.com/slok/brigadeterm/pkg/model/brigade"
)

func TestControllerProjectListPageContext(t *testing.T) {
	start := time.Now().Add(-5 * time.Minute)
	end := time.Now().Add(-4 * time.Minute)

	tests := []struct {
		name      string
		projects  []*brigademodel.Project
		lastBuild *brigademodel.Build
		expCtx    *controller.ProjectListPageContext
	}{
		{
			name: "One project should return one project context.",
			projects: []*brigademodel.Project{
				&brigademodel.Project{
					ID:   "prj1",
					Name: "project-1",
				},
			},
			lastBuild: &brigademodel.Build{
				ID:       "build1",
				Revision: &azurebrigade.Revision{Commit: "1234567890"},
				Worker: &azurebrigade.Worker{
					Status:    azurebrigade.JobSucceeded,
					StartTime: start,
					EndTime:   end,
				},
				Type: "testEvent",
			},
			expCtx: &controller.ProjectListPageContext{
				Projects: []*controller.Project{
					&controller.Project{
						ID:   "prj1",
						Name: "project-1",
						LastBuild: &controller.Build{
							ID:         "build1",
							Version:    "1234567890",
							EventType:  "testEvent",
							Running:    false,
							FinishedOK: true,
							Started:    start,
							Ended:      end,
						},
					},
				},
			},
		},
		{
			name: "Multiple projects should return multiple project context.",
			projects: []*brigademodel.Project{
				&brigademodel.Project{
					ID:   "prj1",
					Name: "project-1",
				},
				&brigademodel.Project{
					ID:   "prj2",
					Name: "project-2",
				},
			},
			lastBuild: &brigademodel.Build{
				ID:       "build1",
				Revision: &azurebrigade.Revision{Commit: "1234567890"},
				Worker: &azurebrigade.Worker{
					Status:    azurebrigade.JobFailed,
					StartTime: start,
					EndTime:   end,
				},
				Type: "testEvent",
			},
			expCtx: &controller.ProjectListPageContext{
				Projects: []*controller.Project{
					&controller.Project{
						ID:   "prj1",
						Name: "project-1",
						LastBuild: &controller.Build{
							ID:         "build1",
							Version:    "1234567890",
							EventType:  "testEvent",
							Running:    false,
							FinishedOK: false,
							Started:    start,
							Ended:      end,
						},
					},
					&controller.Project{
						ID:   "prj2",
						Name: "project-2",
						LastBuild: &controller.Build{
							ID:         "build1",
							Version:    "1234567890",
							EventType:  "testEvent",
							Running:    false,
							FinishedOK: false,
							Started:    start,
							Ended:      end,
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)

			// Mocks.
			mb := &mbrigade.Service{}
			mb.On("GetProjects").Return(test.projects, nil)
			mb.On("GetProjectLastBuild", mock.Anything).Return(test.lastBuild, nil)

			c := controller.NewController(mb)
			ctx := c.ProjectListPageContext()

			assert.Equal(test.expCtx, ctx)
		})
	}
}
