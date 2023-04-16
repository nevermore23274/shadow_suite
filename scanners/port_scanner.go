package port_scanner

import (
	"fmt"
	"net"
	"strconv"
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

func TcpCheck(url string, portEntry string, tcpConn chan string, wg *sync.WaitGroup, statusChan chan string) {

	// Create slice of port(s) given
	portStrings := strings.Split(portEntry, " ")

	// Creating slice for ports
	portSlice := make([]int, len(portStrings))

	// Converting port strings to ints
	for i, s := range portStrings {
		portSlice[i], _ = strconv.Atoi(s)
	}

	// Iterate over given ports and scan
	for i, v := range portSlice {
		tcpConnStr := fmt.Sprintf("%s:%d", url, v)
		conn, err := net.Dial("tcp", tcpConnStr)

		if err != nil {
			statusChan <- fmt.Sprintf("%d closed %s", v, "TCP")
			continue
		}

		conn.Close()

		statusChan <- fmt.Sprintf("%d open %s", v, "TCP")
		time.Sleep(10 * time.Millisecond)
		i++
	}

	defer wg.Done()
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
