package ver

import "runtime"

// Version returns value of 'go version' as a string
func Version() string {
	return runtime.Version()
}
