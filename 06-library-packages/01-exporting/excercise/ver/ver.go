package ver

import "runtime"

// Version prints 'go version' as a string
func Version() string {
	return runtime.Version()
}
