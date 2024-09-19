package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jumpycalm/mview"
)

// End shows the final slide.
func End(nextSlide func()) (title string, info string, content mview.Primitive) {
	textView := mview.NewTextView()
	textView.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	url := "https://github.com/jumpycalm/mview"
	fmt.Fprint(textView, url)
	return "End", "", Center(len(url), 1, textView)
}
