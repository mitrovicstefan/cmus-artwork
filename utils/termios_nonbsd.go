//go:build !darwin && !freebsd && !netbsd && !openbsd && !windows
// +build !darwin,!freebsd,!netbsd,!openbsd,!windows

package utils

import (
	"golang.org/x/sys/unix"
)

const (
	GetTermios = unix.TCGETS
	SetTermios = unix.TCSETS
)
