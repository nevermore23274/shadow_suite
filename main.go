package main

import (
	"network_tool/scanners"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// declare scanButton outside of the main function and export it
var ScannerButton *widget.Button

func main() {
	// Create a new Fyne application and window
	a := app.New()
	mainWindow := a.NewWindow("Scanner")
	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.CenterOnScreen()

	// Create channels for communication between goroutines
	tcpConn := make(chan string)
	statusChan := make(chan string)
	done := make(chan bool)

	// Create a text field for the user to enter the hostname or IP address to scan
	hostnameEntry := widget.NewEntry()
	hostnameEntry.SetText("scanme.nmap.org")

	// Create a label to display TCP connection status
	tcpConnLabel := widget.NewLabel("")

	// Create a waitgroup to keep track of TCP checks
	wg := &sync.WaitGroup{}

	// Create a button to start the TCP scan
	ScannerButton = widget.NewButton("Scanner", func() {
		// Button click event handler goes here
	})

	widgetButton := widget.NewButton("Begin Scan", func() {
		// Update the label to indicate the scan has started
		tcpConnLabel.SetText("")

		// Disable the hostname entry field
		hostnameEntry.Disable()

		// Add a new task to the waitgroup and start a new goroutine to perform the TCP check
		wg.Add(1)
		go func() {
			scanners.TcpCheck(hostnameEntry.Text, tcpConn, wg, statusChan) // use the text entered in the hostnameEntry field
			done <- true
		}()

		// Disable the scan button after it is clicked
		ScannerButton.Disable()
	})
	bottomLeft := container.NewVBox(widgetButton)

	// Create a vertical box to hold the scan button and hostname entry field
	sideBar := container.NewVBox(
		ScannerButton,
	)

	sideBar.Resize(fyne.NewSize(150, 0))

	// Create a scrollable container to hold the TCP connection status label
	tcpConnScroll := container.NewScroll(tcpConnLabel)
	tcpConnScroll.SetMinSize(fyne.NewSize(0, 450)) // Set desired size of container

	// Create a vertical split container to display the sidebar and the main content area
	content := container.NewHSplit(
		sideBar,
		container.NewVBox(
			// create a layout with a vertical layout and add a label to it
			container.New(layout.NewVBoxLayout(), widget.NewLabelWithStyle("Top half", fyne.TextAlignCenter,
				fyne.TextStyle{Bold: true})),
			container.NewMax(tcpConnScroll),
			container.NewGridWithRows(2,
				container.NewMax(bottomLeft),
				container.NewMax(hostnameEntry),
			),
		),
	)

	// Set the size of the top and bottom halves
	content.Resize(fyne.NewSize(0, 400))

	content.Offset = 0.2 // adjust the offset value as desired

	// Set the content of the main window to the horizontal split container
	mainWindow.SetContent(content)

	// Wait for the TCP check to finish and mark the waitgroup as done
	go func() {
		<-done
		wg.Done()
	}()

	// Start a new goroutine to update the TCP connection status label when a message is received on the status channel
	go scanners.UpdateTcpConnLabel(statusChan, tcpConnLabel, done) // pass the statusChan channel

	// Show the main window and start the Fyne event loop
	mainWindow.ShowAndRun()
}
