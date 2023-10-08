package output_test

import (
	"bufio"
	"testing"

	"github.com/cardil/tui/output"
	"github.com/stretchr/testify/assert"
)

func TestNewTestPrinter(t *testing.T) {
	p := output.NewTestPrinterWithAnswers([]string{
		"answer1",
		"answer2",
	})
	reader := p.InOrStdin()
	// read lines from reader
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		p.Printf("line: %s", scanner.Text())
		p.Println()
		p.PrintErrf("errline: %s", scanner.Text())
		p.PrintErrln()
	}

	assert.Equal(t, "line: answer1\nline: answer2\n",
		p.Outputs().Out.String())
	assert.Equal(t, "errline: answer1\nerrline: answer2\n",
		p.Outputs().Err.String())
}
