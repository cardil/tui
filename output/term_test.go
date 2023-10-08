package output_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/cardil/tui/output"
	"github.com/stretchr/testify/assert"
)

func TestIsReaderTerminal(t *testing.T) {
	in := os.NewFile(0, "stdin")
	assert.False(t, output.IsReaderTerminal(in))

	out := os.NewFile(1, "stdout")
	assert.False(t, output.IsWriterTerminal(out))

	var buf bytes.Buffer
	assert.False(t, output.IsWriterTerminal(&buf))
}
