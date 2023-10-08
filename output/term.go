package output

import (
	"io"
	"os"

	"golang.org/x/term"
)

// IsWriterTerminal returns true if the given writer is a real terminal.
func IsWriterTerminal(w io.Writer) bool {
	return isTerminal(w)
}

// IsReaderTerminal returns true if the given reader is a real terminal.
func IsReaderTerminal(r io.Reader) bool {
	return isTerminal(r)
}

func isTerminal(a any) bool {
	f, ok := a.(*os.File)
	return ok && term.IsTerminal(int(f.Fd()))
}
