// +build go1.17

package rawfile

// Force a compile error to "output" a message.
const compileError int = "This package requires manual checks before it can be built on Go 1.17+. See https://github.com/getlantern/netstack/pull/6"
