package output_test

import (
	"os"
	"testing"

	"github.com/cardil/tui/output"
	"github.com/stretchr/testify/assert"
)

func TestOsPrinter(t *testing.T) {
	assert.Equal(t, os.Stdin, output.OsPrinter.InOrStdin())
	assert.Equal(t, os.Stdout, output.OsPrinter.OutOrStdout())
	assert.Equal(t, os.Stderr, output.OsPrinter.ErrOrStderr())
}
