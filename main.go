package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
	"golang.org/x/tools/go/analysis/passes/defers"
	//"bufio"
	//"errors"
	//"os"
)

type Screen struct {
	screen         tcell.Screen
	err            error
	currentMode    string
	x1, y1, x2, y2 int
}

const (
	ModeNormal      = "Normal"
	ModeCommand     = "Command"
	ModeInsert      = "Insert"
	statusBarHeight = 1
)

func main() {
	var s Screen
	s.screen, s.err = tcell.NewScreen()
	if s.err != nil {
		log.Fatalf("%+v", s.err) // Handle error appropriately
	}
	if s.err = s.screen.Init(); s.err != nil {
		log.Fatalf("%+v", s.err) // Handle error appropriately
	}

	s.currentMode = ModeNormal
	s.screen.SetContent()
	// Initialize other editor components here
	drawScreen(&s)

	quit := func() {
		maybePanic := recover()
		s.screen.Fini() // Ensure cleanup on exit
		if maybePanic != nil {
			panic(maybePanic)
		}
	}

	defer quit()
	// Event loop, key handling, etc.

}

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func drawScreen(s *Screen) {
	s.screen.Clear()
	x2, y2 := s.screen.Size()
	drawBox(s, 0, 0, x2, y2, tcell.StyleDefault)
	s.screen.Show()
}

func drawBox(s *Screen, x1, y1, x2, y2 int, style tcell.Style) {
	// Draw top and bottom lines
	for x := x1; x <= x2; x++ {
		s.screen.SetContent(x, y1, tcell.RuneHLine, nil, style)
		s.screen.SetContent(x, y2, tcell.RuneHLine, nil, style)
	}

	// Draw left and right lines
	for y := y1 + 1; y < y2; y++ { // Skip corners
		s.screen.SetContent(x1, y, tcell.RuneVLine, nil, style)
		s.screen.SetContent(x2, y, tcell.RuneVLine, nil, style)
	}

	// Draw corners (optional)
	s.screen.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
	s.screen.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
	s.screen.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
	s.screen.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
}
