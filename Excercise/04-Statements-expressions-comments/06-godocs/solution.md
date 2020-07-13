1. Documentation of `runtime.NumCPU()` 
```
ad@mac:$ go doc -src runtime NumCPU
package runtime // import "runtime"

// NumCPU returns the number of logical CPUs usable by the current process.
//
// The set of available CPUs is checked by querying the operating system
// at process startup. Changes to operating system CPU allocation after
// process startup are not reflected.
```

2. Printing source code of `runtime.NumCPU()`
```
func NumCPU() int {
        return int(ncpu)
}
ad@mac:$
``` 