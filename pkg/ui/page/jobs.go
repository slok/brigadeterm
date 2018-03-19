package page

import (
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
)

// BuildJobList is the page where a build job list will be available.
type BuildJobList struct {
	controller controller.Controller
	router     *Router

	// page layout
	layout tview.Primitive

	// components
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
		SetBorder(true).
		SetTitle("Pipeline timeline")

	b.jobsList = tview.NewTable().
		SetSelectable(true, false)
	b.jobsList.
		SetBorder(true).
		SetTitle("Jobs")

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
	b.fill(ctx)

	// Set key handlers.
	//b.buildsTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	//	switch event.Key() {
	//	case tcell.KeyF5:
	//		// Reload.
	//	case tcell.KeyEsc:
	//		// Back.
	//		b.router.LoadProjectBuildList(projectID)
	//	}
	//	return event
	//})

}

func (b *BuildJobList) fill(ctx *controller.BuildJobListPageContext) {
	b.fillPipelineTimeline(ctx)
}

func (b *BuildJobList) fillPipelineTimeline(ctx *controller.BuildJobListPageContext) {
	// Create our pipeline timeline.
	b.jobsPipeline.Clear()

	// Get timing information, first job, last job and the total runtime time of the run jobs.
	first, _, totalDuration := b.getJobTimingData(ctx)
	stepDuration := int(totalDuration.Seconds()) / pipelineStepTotal

	offset := 1                // Ignore the first cell (is the name of the job).
	rowsBetweenMultiplier := 2 // Left a row between job rows.

	// Create one row for each job.
	for i, job := range ctx.Jobs {
		// Name of job.
		b.jobsPipeline.SetCell(i*rowsBetweenMultiplier, 0, &tview.TableCell{Text: job.Name, Align: tview.AlignCenter, Color: pipelineColor})

		// Get length of pipeline.
		jobDuration := job.Ended.Sub(job.Started)
		pipelineLen := int(jobDuration.Seconds()) / stepDuration

		// Get the start point of the job by getting the start point and
		// calculating the diff until the start of the current job.
		startOffsetTime := job.Started.Sub(first)
		startOffset := int(startOffsetTime.Seconds()) / stepDuration

		for j := startOffset; j < startOffset+pipelineLen; j++ {
			b.jobsPipeline.SetCell(i*rowsBetweenMultiplier, offset+j, &tview.TableCell{Text: pipelineBlockString, BackgroundColor: pipelineColor, Color: pipelineColor})
		}
	}
}

func (b *BuildJobList) getJobTimingData(ctx *controller.BuildJobListPageContext) (first, last time.Time, totalDuration time.Duration) {
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
