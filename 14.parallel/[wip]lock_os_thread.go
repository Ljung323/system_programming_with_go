package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	runtime.LockOSThread() // Lock the main OS thread
	defer runtime.UnlockOSThread()

	gtk.Init(nil)

	// Create a new window
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("My Window")
	window.SetSizeRequest(800, 600)

	// Connect the "destroy" signal to quit the main loop
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Show the window
	window.ShowAll()

	// Run the main event loop
	for {
		gtk.MainIterationDo(false)

		// Do some work here
		fmt.Println("Doing some work...")

		// Sleep for a bit to avoid hogging the CPU
		time.Sleep(100 * time.Millisecond)
	}
}
