package ui

import (
	"fmt"
	"sync"

	"github.com/rivo/tview"

	"github.com/slok/brigadeterm/pkg/controller"
	"github.com/slok/brigadeterm/pkg/ui/page"
)

// Renderer will render windows.
type Renderer interface {
	Render() error
}

// Index will compose index window.
type Index struct {
	app               *tview.Application
	layout            *tview.Flex
	controller        controller.Controller
	router            *page.Router
	registerPagesOnce sync.Once
}

// NewIndex returns a new index renderer.
func NewIndex(controller controller.Controller, app *tview.Application) *Index {
	// TODO: use brigade service.
	i := &Index{
		app:        app,
		controller: controller,
	}

	i.createLayout()
	return i
}

func (i *Index) createPages() *tview.Pages {
	// Create the tui pages.
	pages := tview.NewPages()

	// Create the page router (also creates and registers the pages on the page ui container).
	i.router = page.NewRouter(i.controller, pages)

	return pages
}

func (i *Index) createLayout() {
	// Create a top row to show global information.
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false)

	// Create the pages.
	pages := i.createPages()

	// Create our layout.
	i.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(info, 1, 1, false).
		AddItem(pages, 0, 1, true)

	// Set title.
	info.SetText(fmt.Sprintf("Welcome to brigadeterm %s", "v0.1.0dev"))

}

// Render satisfies Renderer interface.
func (i *Index) Render() error {
	// Load the initial page.
	i.router.LoadAbout()
	//i.router.LoadProjectList()

	// Run
	i.app.SetRoot(i.layout, true)
	return i.app.Run()
}
