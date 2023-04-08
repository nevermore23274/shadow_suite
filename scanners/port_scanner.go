package scanners

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"fyne.io/fyne/v2/widget"
)

/*
// tcpCheck scans ports 1 through 1024 of "scanme.nmap.org".
// It sends a string with the format "{port number} open/closed" to statusChan for each port scanned.
// If an error occurs while attempting to connect to a port, no message is sent to statusChan.
// wg is used to coordinate goroutines
*/

func TcpCheck(url string, tcpConn chan string, wg *sync.WaitGroup, statusChan chan string) {
	defer wg.Done()

	for i := 1; i <= 1024; i++ {
		tcpConnStr := fmt.Sprintf("%s:%d", url, i)
		conn, err := net.Dial("tcp", tcpConnStr)

		// Send nothing to tcpConn channel if the port is closed
		if err != nil {
			statusChan <- fmt.Sprintf("%d closed %s", i, "TCP")
			continue
		}

		conn.Close()

		// Send the open port message back to the main goroutine via the statusChan
		statusChan <- fmt.Sprintf("%d open %s", i, "TCP")

		// Add a delay to slow down the function
		time.Sleep(10 * time.Millisecond)
	}
}

func UpdateTcpConnLabel(statusChan chan string, tcpConnLabel *widget.Label, done chan bool) {
	// Set initial label text to indicate that the scan has started
	tcpConnLabel.SetText("")

	// Loop over messages in the channel until it's closed
	for msg := range statusChan {
		parts := strings.Split(msg, " ") // Split the message into two parts: port number and status

		// If the message doesn't have two parts, skip it
		if len(parts) != 3 {
			continue
		}

		// Get the port number, status, and scan type from the message
		port, status, scanType := parts[0], parts[1], parts[2]

		// Update the label text with the new port, status, and scan type
		tcpConnLabel.SetText(tcpConnLabel.Text + fmt.Sprintf("%s %s (%s)\n", port, status, scanType))
	}

	// When the channel is closed, indicate that the scan is complete
	tcpConnLabel.SetText(tcpConnLabel.Text + "Done.\n")

	// Signal that the function has finished by sending a value to the done channel
	done <- true
}
