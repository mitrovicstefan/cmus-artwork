//go:build darwin || freebsd || openbsd || netbsd
// +build darwin freebsd openbsd netbsd

package utils

import (
	"golang.org/x/sys/unix"
)

const (
	GetTermios = unix.TIOCGETA
	SetTermios = unix.TIOCSETA
)
