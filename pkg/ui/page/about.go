package page

import (
	"fmt"
	"sync"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const (
	textview = `
██████╗ ██████╗ ██╗ ██████╗  █████╗ ██████╗ ███████╗████████╗███████╗██████╗ ███╗   ███╗
██╔══██╗██╔══██╗██║██╔════╝ ██╔══██╗██╔══██╗██╔════╝╚══██╔══╝██╔════╝██╔══██╗████╗ ████║
██████╔╝██████╔╝██║██║  ███╗███████║██║  ██║█████╗     ██║   █████╗  ██████╔╝██╔████╔██║
██╔══██╗██╔══██╗██║██║   ██║██╔══██║██║  ██║██╔══╝     ██║   ██╔══╝  ██╔══██╗██║╚██╔╝██║
██████╔╝██║  ██║██║╚██████╔╝██║  ██║██████╔╝███████╗   ██║   ███████╗██║  ██║██║ ╚═╝ ██║
╚═════╝ ╚═╝  ╚═╝╚═╝ ╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚══════╝   ╚═╝   ╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝
																						
press any key to continue.
`
)

const (
	// AboutPageName is the name that identifies thi About page.
	AboutPageName = "about"
)

// About is the page where the about will reside.
type About struct {
	router *Router

	// page layout
	layout tview.Primitive

	// components
	aboutText *tview.TextView

	registerPageOnce sync.Once
}

// NewAbout returns a new About.
func NewAbout(router *Router) *About {
	a := &About{
		router: router,
	}
	a.createComponents()
	return a
}

// createComponents will create all the layout components.
func (a *About) createComponents() {
	a.aboutText = tview.NewTextView().
		SetTextColor(tcell.ColorYellow).
		SetScrollable(false)

	// Create the layout.
	a.layout = tview.NewFlex().
		AddItem(a.aboutText, 0, 1, true)

	// Set handler.
	a.aboutText.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Always go to first page on any key.
		a.router.LoadProjectList()
		return event
	})
}

// Register satisfies Page interface.
func (a *About) Register(pages *tview.Pages) {
	a.registerPageOnce.Do(func() {
		pages.AddPage(AboutPageName, a.layout, true, false)
	})
}

// BeforeLoad satisfies Page interface.
func (a *About) BeforeLoad() {
}

// Refresh will refresh all the page data.
func (a *About) Refresh() {
	a.aboutText.Clear()
	fmt.Fprintf(a.aboutText, textview)
}
