package page

import (
	"github.com/rivo/tview"
	"github.com/slok/brigadeterm/pkg/controller"
)

// Router knows how to route the ui from one page to another.
type Router struct {
	pages                *tview.Pages
	projectListPage      *ProjectList
	projectBuildListPage *ProjectBuildList
	buildJobListPage     *BuildJobList
	jobLogPage           *JobLog
	app                  *tview.Application
}

// NewRouter returns a new router.
func NewRouter(app *tview.Application, controller controller.Controller, pages *tview.Pages) *Router {
	r := &Router{
		pages: pages,
		app:   app,
	}

	// Create the pages.
	r.projectListPage = NewProjectList(controller, r)
	r.projectBuildListPage = NewProjectBuildList(controller, r)
	r.buildJobListPage = NewBuildJobList(controller, r)
	r.jobLogPage = NewJobLog(controller, r)

	// Register our pages on the app pages container.
	r.register()

	return r
}

// Register will register the pages on the ui
func (r *Router) register() {
	pages := []Page{
		r.projectListPage,
		r.projectBuildListPage,
		r.buildJobListPage,
		r.jobLogPage,
	}

	// Register all the pages on the ui.
	for _, page := range pages {
		page.Register(r.pages)
	}
}

// LoadProjectList will set the ui on the project list.
func (r *Router) LoadProjectList() {
	r.projectListPage.BeforeLoad()
	r.projectListPage.Refresh()
	r.pages.SwitchToPage(ProjectListPageName)
}

// LoadProjectBuildList will set the ui on the project build list.
func (r *Router) LoadProjectBuildList(projectID string) {
	r.projectBuildListPage.BeforeLoad()
	r.projectBuildListPage.Refresh(projectID)
	r.pages.SwitchToPage(ProjectBuildListPageName)
}

// LoadBuildJobList will set the ui on the build job list.
func (r *Router) LoadBuildJobList(projectID, buildID string) {
	r.buildJobListPage.BeforeLoad()
	r.buildJobListPage.Refresh(projectID, buildID)
	r.pages.SwitchToPage(BuildJobListPageName)
}

// LoadJobLog will set the ui on the build job log.
func (r *Router) LoadJobLog(projectID, buildID, jobID string) {
	r.jobLogPage.BeforeLoad()
	r.jobLogPage.Refresh(projectID, buildID, jobID)
	r.pages.SwitchToPage(JobLogPageName)
}

// Exit will terminate everything.
func (r *Router) Exit() {
	r.app.Stop()
}
