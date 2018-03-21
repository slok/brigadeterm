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

func TestControllerProjectBuildListPageContext(t *testing.T) {
	start := time.Now().Add(-5 * time.Minute)
	end := time.Now().Add(-4 * time.Minute)

	tests := []struct {
		name    string
		project *brigademodel.Project
		builds  []*brigademodel.Build
		expCtx  *controller.ProjectBuildListPageContext
	}{
		{
			name: "Multiple builds should return multiple builds context.",
			project: &brigademodel.Project{
				ID:         "prj1",
				Name:       "project-1",
				Kubernetes: azurebrigade.Kubernetes{Namespace: "test"},
				Repo:       azurebrigade.Repo{CloneURL: "git@github.com:slok/brigadeterm"},
			},
			builds: []*brigademodel.Build{
				&brigademodel.Build{
					ID:       "build1",
					Revision: &azurebrigade.Revision{Commit: "1234567890"},
					Worker: &azurebrigade.Worker{
						Status:    azurebrigade.JobFailed,
						StartTime: start,
						EndTime:   end,
					},
					Type: "testEvent",
				},
				&brigademodel.Build{
					ID:       "build2",
					Revision: &azurebrigade.Revision{Commit: "1234567890"},
					Worker: &azurebrigade.Worker{
						Status:    azurebrigade.JobSucceeded,
						StartTime: start,
						EndTime:   end,
					},
					Type: "testEvent",
				},
				&brigademodel.Build{
					ID:       "build3",
					Revision: &azurebrigade.Revision{Commit: "1234567890"},
					Worker: &azurebrigade.Worker{
						Status:    azurebrigade.JobRunning,
						StartTime: start,
					},
					Type: "testEvent",
				},
			},
			expCtx: &controller.ProjectBuildListPageContext{
				ProjectName: "project-1",
				ProjectNS:   "test",
				ProjectURL:  "git@github.com:slok/brigadeterm",
				Builds: []*controller.Build{
					&controller.Build{
						ID:         "build1",
						Version:    "1234567890",
						Running:    false,
						FinishedOK: false,
						EventType:  "testEvent",
						Started:    start,
						Ended:      end,
					},
					&controller.Build{
						ID:         "build2",
						Version:    "1234567890",
						Running:    false,
						FinishedOK: true,
						EventType:  "testEvent",
						Started:    start,
						Ended:      end,
					},
					&controller.Build{
						ID:         "build3",
						Version:    "1234567890",
						Running:    true,
						FinishedOK: false,
						EventType:  "testEvent",
						Started:    start,
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
			mb.On("GetProject", mock.Anything).Return(test.project, nil)
			mb.On("GetProjectBuilds", mock.Anything).Return(test.builds, nil)

			c := controller.NewController(mb)
			ctx := c.ProjectBuildListPageContext("whatever")

			assert.Equal(test.expCtx, ctx)
		})
	}
}
