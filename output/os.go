package output

import (
	"io"
	"os"
)

// OsInOut is an InputOutput implementation that uses the standard input,
// output, and error streams.
type OsInOut struct{}

// OsPrinter is the default printer for the CLI. It uses the standard input,
// output, and error streams.
var OsPrinter = stdPrinter{OsInOut{}} //nolint:gochecknoglobals

func (o OsInOut) InOrStdin() io.Reader {
	return os.Stdin
}

func (o OsInOut) OutOrStdout() io.Writer {
	return os.Stdout
}

func (o OsInOut) ErrOrStderr() io.Writer {
	return os.Stderr
}

var _ InputOutput = OsInOut{}
