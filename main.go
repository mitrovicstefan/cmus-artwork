package main

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
  "os/exec"
  "golang.org/x/sys/unix"
  "github.com/mitrovicstefan/cmus-artwork/utils"
)

func main() {
	// Channel to receive input characters
	inputChannel := make(chan rune)
  // Store the original terminal state
  fd := int(os.Stdin.Fd())
  oldState, err := unix.IoctlGetTermios(fd, unix.TCGETS)
  if err != nil {
    fmt.Println("Error getting terminal attributes:", err)
    return
  }

  // Ensure terminal state is restored on exit
	defer unix.IoctlSetTermios(fd, unix.TCSETS, oldState)

	// Start terminal update goroutine
	go utils.UpdateTerminal()

	// Start input handling goroutine
	go utils.HandleInput(inputChannel)

	// Setup signal handling to exit gracefully
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)

	// Main loop to process input
	for {
		select {
		case input := <-inputChannel:
      if input == rune('j') {
        cmd := exec.Command("cmus-remote", "--pause")
        cmd.Output()
      }
      if input == rune('h') {
        cmd := exec.Command("cmus-remote", "--prev")
        cmd.Output()
      }
      if input == rune('l') {
        cmd := exec.Command("cmus-remote", "--next")
        cmd.Output()
      }
		case <-sigChannel:
			fmt.Printf("Thank you!\n")
			return
		}
	}
}

