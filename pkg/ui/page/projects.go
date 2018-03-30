package page

import (
	"fmt"
	"sync"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/slok/brigadeterm/pkg/controller"
)

const (
	// ProjectListPageName is the name that identifies thi projectList page.
	ProjectListPageName = "projectlist"

	projectListUsage = `[yellow](F5) [white]Reload    [yellow](ctrl+Q) [white]Quit`
)

// ProjectList is the main page where the project list will be available.
type ProjectList struct {
	controller controller.Controller
	router     *Router

	// page layout
	layout tview.Primitive

	// components
	projectsTable *tview.Table
	usage         *tview.TextView

	registerPageOnce sync.Once
}

// NewProjectList returns a new project list.
func NewProjectList(controller controller.Controller, router *Router) *ProjectList {
	p := &ProjectList{
		controller: controller,
		router:     router,
	}
	p.createComponents()
	return p
}

// Register satisfies Page interface.
func (p *ProjectList) Register(pages *tview.Pages) {
	p.registerPageOnce.Do(func() {
		pages.AddPage(ProjectListPageName, p.layout, true, false)
	})
}

// BeforeLoad satisfies Page interface.
func (p *ProjectList) BeforeLoad() {
}

// Refresh will refresh all the page data.
func (p *ProjectList) Refresh() {
	ctx := p.controller.ProjectListPageContext()
	// TODO: check error.
	p.fill(ctx)

	// Set key handlers.
	p.projectsTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyF5:
			// Reload.
			p.router.LoadProjectList()
		case tcell.KeyCtrlQ:
			p.router.Exit()
		}
		return event
	})

}

// createComponents will create all the layout components.
func (p *ProjectList) createComponents() {
	// Set up columns
	p.projectsTable = tview.NewTable().
		SetSelectable(true, false)
	p.projectsTable.
		SetBorder(true).
		SetTitle("Projects")

	// Usage.
	p.usage = tview.NewTextView().
		SetDynamicColors(true)

	// Create the layout.
	p.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(p.projectsTable, 0, 1, true).
		AddItem(p.usage, 1, 1, false)
}

func (p *ProjectList) fill(ctx *controller.ProjectListPageContext) {
	p.fillUsage()
	p.fillProjectList(ctx)
}

func (p *ProjectList) fillUsage() {
	p.usage.Clear()
	p.usage.SetText(projectListUsage)
}

func (p *ProjectList) fillProjectList(ctx *controller.ProjectListPageContext) {
	// Clear other widgets.
	p.projectsTable.Clear()

	// Set header.
	p.projectsTable.SetCell(0, 0, &tview.TableCell{Align: tview.AlignCenter, Color: tcell.ColorYellow})
	p.projectsTable.SetCell(0, 1, &tview.TableCell{Text: "Name", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	p.projectsTable.SetCell(0, 2, &tview.TableCell{Text: "Last build type", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	p.projectsTable.SetCell(0, 3, &tview.TableCell{Text: "Last build version", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	p.projectsTable.SetCell(0, 4, &tview.TableCell{Text: "Last build time", Align: tview.AlignCenter, Color: tcell.ColorYellow})

	projectNameIDIndex := map[string]string{}

	// Set body.
	rowPosition := 1
	for _, project := range ctx.Projects {
		if project == nil {
			continue
		}

		var event string
		var version string
		var since time.Duration
		color := unknownColor
		icon := unknownIcon

		if project.LastBuild != nil {
			color = getColorFromState(project.LastBuild.State)
			icon = getIconFromState(project.LastBuild.State)

			// Calculate lastbuild data.
			event = project.LastBuild.EventType
			version = project.LastBuild.Version
			since = time.Since(project.LastBuild.Started).Truncate(time.Second * 1)
		}

		// Set the index so we can get the project ID on selection.
		projectNameIDIndex[project.Name] = project.ID

		p.projectsTable.SetCell(rowPosition, 0, &tview.TableCell{Text: icon, Align: tview.AlignLeft, Color: color})
		p.projectsTable.SetCell(rowPosition, 1, &tview.TableCell{Text: project.Name, Align: tview.AlignLeft, Color: color})
		p.projectsTable.SetCell(rowPosition, 2, &tview.TableCell{Text: event, Align: tview.AlignLeft, Color: color})
		p.projectsTable.SetCell(rowPosition, 3, &tview.TableCell{Text: version, Align: tview.AlignLeft, Color: color})
		p.projectsTable.SetCell(rowPosition, 4, &tview.TableCell{Text: fmt.Sprintf("%v ago", since), Align: tview.AlignLeft, Color: color})

		rowPosition++
	}

	// Set selectable to call our jobs.
	p.projectsTable.SetSelectedFunc(func(row, column int) {
		// If the row is the header then don't do anything.
		if row > 0 {
			// Get project ID cell and from commit the build ID.
			cell := p.projectsTable.GetCell(row, 1)
			projectID := projectNameIDIndex[cell.Text]
			// Load build list page.
			p.router.LoadProjectBuildList(projectID)
		}
	})
}
