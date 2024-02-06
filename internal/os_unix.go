package internal

import (
	"syscall"
)

func osShouldRetry(err error) bool {
	return err == syscall.ECONNRESET || err == syscall.ECONNREFUSED
}
