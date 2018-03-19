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
	aboutPage            *About
}

// NewRouter returns a new router.
func NewRouter(controller controller.Controller, pages *tview.Pages) *Router {
	r := &Router{
		pages: pages,
	}

	// Create the pages.
	r.projectListPage = NewProjectList(controller, r)
	r.projectBuildListPage = NewProjectBuildList(controller, r)
	r.buildJobListPage = NewBuildJobList(controller, r)
	r.aboutPage = NewAbout(r)

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
		r.aboutPage,
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
func (r *Router) LoadBuildJobList(projectID string, buildID string) {
	r.buildJobListPage.BeforeLoad()
	r.buildJobListPage.Refresh(projectID, buildID)
	r.pages.SwitchToPage(BuildJobListPageName)
}

// LoadAbout will set the ui on the about list
func (r *Router) LoadAbout() {
	r.aboutPage.BeforeLoad()
	r.aboutPage.Refresh()
	r.pages.SwitchToPage(AboutPageName)
}
