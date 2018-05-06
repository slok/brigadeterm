package page

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/slok/brigadeterm/pkg/controller"
)

const (
	// ProjectListPageName is the name that identifies thi projectList page.
	ProjectListPageName = "projectlist"

	projectListUsage = `[yellow](F5) [white]Reload    [yellow](/) [white]Filter    [yellow](Q) [white]Quit`
)

var projectListFilter string

// ProjectList is the main page where the project list will be available.
type ProjectList struct {
	controller controller.Controller
	router     *Router

	// page layout
	layout tview.Primitive

	// components
	projectsTable    *tview.Table
	usage            *tview.TextView
	filterInputField *tview.InputField

	registerPageOnce sync.Once
}

func (p *ProjectList) focusFilterForm() {
	p.filterInputField.SetLabelColor(tcell.ColorYellow)
	p.router.app.SetFocus(p.filterInputField)
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
		// Reload.
		case tcell.KeyF5:
			p.router.LoadProjectList()
		// Regular keys handling:
		case tcell.KeyRune:
			switch event.Rune() {
			// filter
			case '/':
				p.focusFilterForm()
			// Reload.
			case 'r', 'R':
				p.router.LoadProjectList()
			// Exit
			case 'q', 'Q':
				p.router.Exit()
			}
		}
		return event
	})

}

// createComponents will create all the layout components.
func (p *ProjectList) createComponents() {
	p.filterInputField = tview.NewInputField().
		SetLabel("Filter: ")
	p.filterInputField.
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetLabelColor(tcell.ColorBlack).
		SetDoneFunc(func(key tcell.Key) {
			term := p.filterInputField.GetText()
			if term == "" {
				p.filterInputField.SetLabelColor(tcell.ColorBlack)
			} else {
				p.filterInputField.SetLabelColor(tcell.ColorYellow)
			}
			p.setFilter(term)
			p.filter()
			p.router.app.SetFocus(p.projectsTable)
		})
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
		// AddItem(p.projectsTable, 0, 1, true).
		AddItem(p.projectsTable, 0, 1, true).
		AddItem(p.usage, 1, 1, false).
		AddItem(p.filterInputField, 1, 1, true)
}

func (p *ProjectList) fill(ctx *controller.ProjectListPageContext) {
	p.fillUsage()
	p.fillProjectList(ctx)
}

func (p *ProjectList) fillUsage() {
	p.usage.Clear()
	p.usage.SetText(projectListUsage)
}

func (p *ProjectList) setFilter(term string) {
	projectListFilter = term
}

func (p *ProjectList) filterSetItemVisibility(rowIndex int, visibility bool) {
	var colIndex int
	var colCount int
	colCount = p.projectsTable.GetColumnCount()
	for colIndex = 0; colIndex < colCount; colIndex++ {
		myTCell := p.projectsTable.GetCell(rowIndex, colIndex)
		if visibility == true {
			myTCell.SetSelectable(true)
		} else {
			myTCell.SetSelectable(false)
		}
		if colIndex < 5 { // for everything, except the "icons"
			if visibility == true {
				myTCell.SetTextColor(tcell.ColorGray)
			} else {
				myTCell.SetTextColor(tcell.ColorBlack)
			}
		} else { // only for the "icons"
			if visibility == true {
				myTCell.SetBackgroundColor(tcell.ColorGray)
			} else {
				myTCell.SetBackgroundColor(tcell.ColorBlack)
			}
		}
	}
}

func (p *ProjectList) filter() {
	var projectsCount int
	var rowIndex int
	projectsCount = p.projectsTable.GetRowCount()
	for rowIndex = 1; rowIndex < projectsCount; rowIndex++ {
		tableCell := p.projectsTable.GetCell(rowIndex, 1)
		projectName := tableCell.Text
		// hide the current row, if it does not match the
		// projectListFilter
		if projectListFilter != "" && strings.Contains(strings.ToLower(projectName), projectListFilter) == false {
			p.filterSetItemVisibility(rowIndex, false)
		} else { // show the current row
			p.filterSetItemVisibility(rowIndex, true)
		}

	}
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

		var lastBuild *controller.Build
		if len(project.LastBuilds) >= 1 {
			lastBuild = project.LastBuilds[0]
		}

		if lastBuild != nil {
			color = getColorFromState(lastBuild.State)
			icon = getIconFromState(lastBuild.State)

			// Calculate lastbuild data.
			event = lastBuild.EventType
			version = lastBuild.Version
			since = time.Since(lastBuild.Started).Truncate(time.Second * 1)
		}

		// Set the index so we can get the project ID on selection.
		projectNameIDIndex[project.Name] = project.ID

		p.projectsTable.SetCell(rowPosition, 0, &tview.TableCell{Text: icon, Align: tview.AlignLeft, Color: color})
		p.projectsTable.SetCell(rowPosition, 1, &tview.TableCell{Text: project.Name, Align: tview.AlignLeft, Color: color})
		p.projectsTable.SetCell(rowPosition, 2, &tview.TableCell{Text: event, Align: tview.AlignLeft, Color: color})
		p.projectsTable.SetCell(rowPosition, 3, &tview.TableCell{Text: version, Align: tview.AlignLeft, Color: color})
		p.projectsTable.SetCell(rowPosition, 4, &tview.TableCell{Text: fmt.Sprintf("%v ago", since), Align: tview.AlignLeft, Color: color})

		// Add the last build status
		columnPosition := 5
		for _, b := range project.LastBuilds {
			color := unknownColor
			icon := unknownIcon
			if lastBuild != nil {
				color = getColorFromState(b.State)
				icon = getIconFromState(b.State)
			}
			p.projectsTable.SetCell(rowPosition, columnPosition, &tview.TableCell{Text: fmt.Sprintf(" %s ", icon), Align: tview.AlignCenter, BackgroundColor: color, Color: tcell.ColorBlack})
			columnPosition++
		}
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
	p.filter()
}
