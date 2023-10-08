package output

import (
	"io"
)

// InputOutput is an interface for types that can provide input and output
// and error streams. It is compatible with Cobra functions, so that the
// cobra.Command can be passed an InputOutput implementation.
type InputOutput interface {
	InOrStdin() io.Reader
	OutOrStdout() io.Writer
	ErrOrStderr() io.Writer
}
