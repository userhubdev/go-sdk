//go:build !unix

package internal

func osShouldRetry(err error) bool {
	return false
}
