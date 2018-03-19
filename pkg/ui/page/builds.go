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
	projectInfoFMT = `[yellow]Project: [white]%s
[yellow]URL: [white]%s
[yellow]Namespace: [white]%s`
	projectBuildListUsage = `[yellow](F5) [white]Reload	[yellow](ESC) [white]Back`
)

const (
	// ProjectBuildListPageName is the name that identifies thi ProjectBuildList page.
	ProjectBuildListPageName = "projectbuildlist"
)

// ProjectBuildList is the page where a projects build list will be available.
type ProjectBuildList struct {
	controller controller.Controller
	router     *Router

	// page layout
	layout tview.Primitive

	// components
	projectInfo *tview.TextView
	buildsTable *tview.Table
	usage       *tview.TextView

	registerPageOnce sync.Once
}

// NewProjectBuildList returns a new ProjectBuildList.
func NewProjectBuildList(controller controller.Controller, router *Router) *ProjectBuildList {
	p := &ProjectBuildList{
		controller: controller,
		router:     router,
	}
	p.createComponents()
	return p
}

// createComponents will create all the layout components.
func (p *ProjectBuildList) createComponents() {
	p.projectInfo = tview.NewTextView().
		SetDynamicColors(true)
	p.projectInfo.SetBorder(true).
		SetBorderColor(tcell.ColorYellow)

	p.buildsTable = tview.NewTable().
		SetSelectable(true, false)
	p.buildsTable.
		SetBorder(true).
		SetTitle("Builds")

	p.usage = tview.NewTextView().
		SetDynamicColors(true)

	// Create the layout.
	p.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(p.projectInfo, 0, 1, false).
		AddItem(p.buildsTable, 0, 6, true).
		AddItem(p.usage, 1, 1, false)

}

// Register satisfies Page interface.
func (p *ProjectBuildList) Register(pages *tview.Pages) {
	p.registerPageOnce.Do(func() {
		pages.AddPage(ProjectBuildListPageName, p.layout, true, false)
	})
}

// BeforeLoad satisfies Page interface.
func (p *ProjectBuildList) BeforeLoad() {
}

// Refresh will refresh all the page data.
func (p *ProjectBuildList) Refresh(projectID string) {
	ctx := p.controller.ProjectBuildListPageContext(projectID)
	// TODO: check error.
	p.fill(projectID, ctx)

	// Set key handlers.
	p.buildsTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyF5:
			// Reload.
			p.router.LoadProjectBuildList(projectID)
		case tcell.KeyEsc:
			// Back.
			p.router.LoadProjectList()
		}
		return event
	})

}

func (p *ProjectBuildList) fill(projectID string, ctx *controller.ProjectBuildListPageContext) {
	// Fill the project information.
	p.projectInfo.Clear()
	p.projectInfo.SetText(fmt.Sprintf(projectInfoFMT, ctx.ProjectName, ctx.ProjectURL, ctx.ProjectNS))

	// Fill usage (not required).
	p.usage.Clear()
	p.usage.SetText(projectBuildListUsage)

	// Fill the build table.
	p.buildsTable.Clear()

	// Set header.
	p.buildsTable.SetCell(0, 0, &tview.TableCell{Text: "Event type", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	p.buildsTable.SetCell(0, 1, &tview.TableCell{Text: "Version", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	p.buildsTable.SetCell(0, 2, &tview.TableCell{Text: "ID", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	p.buildsTable.SetCell(0, 3, &tview.TableCell{Text: "End", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	p.buildsTable.SetCell(0, 4, &tview.TableCell{Text: "Duration", Align: tview.AlignCenter, Color: tcell.ColorYellow})

	// TODO order by time.
	rowPosition := 1
	for _, build := range ctx.Builds {
		// Select row color.
		color := tcell.ColorWhite
		if !build.Running {
			if build.FinishedOK {
				color = tcell.ColorGreen
			} else {
				color = tcell.ColorRed
			}
		}
		// Fill table.
		p.buildsTable.SetCell(rowPosition, 0, &tview.TableCell{Text: build.EventType, Align: tview.AlignLeft, Color: color})
		p.buildsTable.SetCell(rowPosition, 1, &tview.TableCell{Text: build.Version, Align: tview.AlignLeft, Color: color})
		p.buildsTable.SetCell(rowPosition, 2, &tview.TableCell{Text: build.ID, Align: tview.AlignLeft, Color: color})
		if !build.Running {
			timeAgo := time.Now().Sub(build.Ended)
			p.buildsTable.SetCell(rowPosition, 3, &tview.TableCell{Text: fmt.Sprintf("%v ago", timeAgo), Align: tview.AlignLeft, Color: color})
			duration := build.Ended.Sub(build.Started)
			p.buildsTable.SetCell(rowPosition, 4, &tview.TableCell{Text: fmt.Sprintf("%v", duration), Align: tview.AlignLeft, Color: color})
		}
		rowPosition++
	}

	// Set selectable to call our jobs.
	p.buildsTable.SetSelectedFunc(func(row, column int) {
		buildID := p.buildsTable.GetCell(row, 2).Text

		// Load build job list page.
		p.router.LoadBuildJobList(projectID, buildID)
	})
}
