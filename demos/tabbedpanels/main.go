// Demo code for the TabbedPanels primitive.
package main

import (
	"github.com/jumpycalm/mview"
)

func main() {
	app := mview.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	tp := mview.NewTabbedPanels()

	// Tab 1 has a textview
	tv := mview.NewTextView()
	tv.SetBorder(true)
	tv.SetTitle("TextView title")
	tv.SetText("TextView text")
	tp.AddTab("tab1", "Tab TextView", tv)

	// Tab 2 has a form
	fm := mview.NewForm()
	fm.SetBorder(true)
	fm.SetTitle("This is the form from the 2nd tab")
	fm.AddInputField("Input", "", 30, nil, nil)
	fm.AddButton("Enter", func() {
		// Don't handle enter
	})
	fm.AddButton("Quit", func() {
		app.Stop()
	})
	fm.SetCancelFunc(func() {
		app.Stop()
	})
	tp.AddTab("tab2", "Tab Form", fm)

	// Tab 3 has a stacked tabs
	tp3 := mview.NewTabbedPanels()
	tp3.AddTab("tab1", "Tab TextView", tv)
	tp3.AddTab("tab2", "Tab Form", fm)
	tp.AddTab("tab3", "Tab stacked", tp3)

	app.SetRoot(tp, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
