package utils

import (
  "fmt"
  "os"
  "golang.org/x/sys/unix"
)

func HandleInput(inputChannel chan rune) {
	fd := int(os.Stdin.Fd())
	oldState, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		fmt.Println("Error getting terminal attributes:", err)
		return
	}

	newState := *oldState
	newState.Lflag &^= unix.ICANON | unix.ECHO
	if err := unix.IoctlSetTermios(fd, unix.TCSETS, &newState); err != nil {
		fmt.Println("Error setting terminal attributes:", err)
		return
	}

	for {
		var buf [1]byte
		_, err := os.Stdin.Read(buf[:])
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		inputChannel <- rune(buf[0])
	}
}
