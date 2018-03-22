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
	// BuildJobListPageName is the name that identifies the BuildJobList page.
	BuildJobListPageName = "buildJoblist"
)

const (
	pipelineBlockString = `█`
	pipelineStepTotal   = 50
	pipelineColor       = tcell.ColorYellow
	pipelineTitle       = "Pipeline timeline (total duration: %s)"

	buildJobListUsage = `[yellow](F5) [white]Reload	[yellow](ESC/Del) [white]Back [yellow](F1) [white]Home`
)

// BuildJobList is the page where a build job list will be available.
type BuildJobList struct {
	controller controller.Controller
	router     *Router

	// page layout.
	layout tview.Primitive

	// components.
	jobsPipeline *tview.Table
	jobsList     *tview.Table
	usage        *tview.TextView

	registerPageOnce sync.Once
}

// NewBuildJobList returns a new BuildJobList.
func NewBuildJobList(controller controller.Controller, router *Router) *BuildJobList {
	b := &BuildJobList{
		controller: controller,
		router:     router,
	}
	b.createComponents()
	return b
}

// createComponents will create all the layout components.
func (b *BuildJobList) createComponents() {
	b.jobsPipeline = tview.NewTable().
		SetBordersColor(pipelineColor)
	b.jobsPipeline.
		SetTitle(fmt.Sprintf(pipelineTitle, time.Millisecond*0)).
		SetBorder(true)

	// Create the job layout (jobs + log).
	b.jobsList = tview.NewTable().
		SetSelectable(true, false)
	b.jobsList.
		SetBorder(true).
		SetTitle("Jobs")

	// Usage.
	b.usage = tview.NewTextView().
		SetDynamicColors(true)

	// Create the layout.
	b.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(b.jobsPipeline, 0, 2, false).
		AddItem(b.jobsList, 0, 6, true).
		AddItem(b.usage, 1, 1, false)
}

// Register satisfies Page interface.
func (b *BuildJobList) Register(pages *tview.Pages) {
	b.registerPageOnce.Do(func() {
		pages.AddPage(BuildJobListPageName, b.layout, true, false)
	})
}

// BeforeLoad satisfies Page interface.
func (b *BuildJobList) BeforeLoad() {
}

// Refresh will refresh all the page data.
func (b *BuildJobList) Refresh(projectID, buildID string) {
	ctx := b.controller.BuildJobListPageContext(buildID)
	// TODO: check error.
	b.fill(projectID, buildID, ctx)

	// Set key handlers.
	b.jobsList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyF5:
			// Reload.
			b.router.LoadBuildJobList(projectID, buildID)
		case tcell.KeyEsc, tcell.KeyDelete:
			// Back.
			b.router.LoadProjectBuildList(projectID)
		case tcell.KeyF1:
			// Home.
			b.router.LoadProjectList()
		}
		return event
	})

}

func (b *BuildJobList) fill(projectID, buildID string, ctx *controller.BuildJobListPageContext) {
	b.fillUsage()
	b.fillPipelineTimeline(ctx)
	b.fillJobsList(projectID, buildID, ctx)
}

func (b *BuildJobList) fillUsage() {
	b.usage.Clear()
	b.usage.SetText(buildJobListUsage)
}

func (b *BuildJobList) fillPipelineTimeline(ctx *controller.BuildJobListPageContext) {
	// Create our pipeline timeline.
	b.jobsPipeline.Clear()

	// Get timing information, first job, last job and the total runtime time of the run jobs.
	first, _, totalDuration := b.getJobTimingData(ctx)
	stepDuration := int(totalDuration.Nanoseconds()) / pipelineStepTotal

	// If there is no duration then don't fill the pipeline
	if stepDuration == 0 {
		return
	}

	offset := 1                // Ignore the first cell (is the name of the job).
	rowsBetweenMultiplier := 2 // Left a row between job rows.

	// Create one row for each job.
	for i, job := range ctx.Jobs {
		// Name of job.
		b.jobsPipeline.SetCell(i*rowsBetweenMultiplier, 0, &tview.TableCell{Text: job.Name, Align: tview.AlignLeft, Color: pipelineColor})

		// Get length of pipeline.
		jobDuration := job.Ended.Sub(job.Started)
		pipelineLen := int(jobDuration.Nanoseconds()) / stepDuration

		// Get the start point of the job by getting the start point and
		// calculating the diff until the start of the current job.
		startOffsetTime := job.Started.Sub(first)
		startOffset := int(startOffsetTime.Nanoseconds()) / stepDuration

		for j := startOffset; j < startOffset+pipelineLen; j++ {
			b.jobsPipeline.SetCell(i*rowsBetweenMultiplier, offset+j, &tview.TableCell{Text: pipelineBlockString, BackgroundColor: pipelineColor, Color: pipelineColor})
		}
	}

	// Set title name:
	b.jobsPipeline.SetTitle(fmt.Sprintf(pipelineTitle, totalDuration.Truncate(1*time.Second)))
}

func (b *BuildJobList) getJobTimingData(ctx *controller.BuildJobListPageContext) (first, last time.Time, totalDuration time.Duration) {
	if len(ctx.Jobs) < 1 {
		return time.Time{}, time.Time{}, 0
	}
	first = ctx.Jobs[0].Started
	last = ctx.Jobs[0].Ended

	for _, job := range ctx.Jobs[1:] {
		// If running is not count.
		if job.Running {
			continue
		}
		if job.Started.Before(first) {
			first = job.Started
		}
		if job.Ended.After(last) {
			last = job.Ended
		}
	}

	return first, last, last.Sub(first)
}

func (b *BuildJobList) fillJobsList(projectID, buildID string, ctx *controller.BuildJobListPageContext) {
	b.jobsList.Clear()

	// Set header.
	b.jobsList.SetCell(0, 0, &tview.TableCell{Align: tview.AlignCenter, Color: tcell.ColorYellow})
	b.jobsList.SetCell(0, 1, &tview.TableCell{Text: "Name", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	b.jobsList.SetCell(0, 2, &tview.TableCell{Text: "Image", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	b.jobsList.SetCell(0, 3, &tview.TableCell{Text: "ID", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	b.jobsList.SetCell(0, 4, &tview.TableCell{Text: "Started", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	b.jobsList.SetCell(0, 5, &tview.TableCell{Text: "Duration", Align: tview.AlignCenter, Color: tcell.ColorYellow})

	// TODO order by time.
	rowPosition := 1
	for _, job := range ctx.Jobs {
		// Select row color and symbol.
		symbol := runningSymbol
		color := tcell.ColorWhite
		if !job.Running {
			if job.FinishedOK {
				symbol = okSymbol
				color = tcell.ColorGreen
			} else {
				symbol = failedSymbol
				color = tcell.ColorRed
			}
		}
		// Fill table.
		b.jobsList.SetCell(rowPosition, 0, &tview.TableCell{Text: symbol, Align: tview.AlignLeft, Color: color})
		b.jobsList.SetCell(rowPosition, 1, &tview.TableCell{Text: job.Name, Align: tview.AlignLeft, Color: color})
		b.jobsList.SetCell(rowPosition, 2, &tview.TableCell{Text: job.Image, Align: tview.AlignLeft, Color: color})
		b.jobsList.SetCell(rowPosition, 3, &tview.TableCell{Text: job.ID, Align: tview.AlignLeft, Color: color})
		if !job.Running {
			timeAgo := time.Since(job.Ended).Truncate(time.Second * 1)
			b.jobsList.SetCell(rowPosition, 4, &tview.TableCell{Text: fmt.Sprintf("%v ago", timeAgo), Align: tview.AlignLeft, Color: color})
			duration := job.Ended.Sub(job.Started).Truncate(time.Second * 1)
			b.jobsList.SetCell(rowPosition, 5, &tview.TableCell{Text: fmt.Sprintf("%v", duration), Align: tview.AlignLeft, Color: color})
		}
		rowPosition++
	}

	// Set selectable to call our jobs.
	b.jobsList.SetSelectedFunc(func(row, column int) {
		jobID := b.jobsList.GetCell(row, 3).Text
		// Load log page
		b.router.LoadJobLog(projectID, buildID, jobID)
	})
}
