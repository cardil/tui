package output

import (
	"bytes"
	"io"
	"strings"
)

// NewTestPrinter returns a TestPrinter with a buffer for input.
func NewTestPrinter() TestPrinter {
	buf := bytes.NewBufferString("")
	return NewTestPrinterWithInput(buf)
}

// NewTestPrinterWithInput returns a TestPrinter with the given input.
func NewTestPrinterWithInput(input io.Reader) TestPrinter {
	return TestPrinter{
		stdPrinter{
			testInOut{
				in: input,
				TestOutputs: TestOutputs{
					Out: new(bytes.Buffer),
					Err: new(bytes.Buffer),
				},
			},
		},
	}
}

// NewTestPrinterWithAnswers returns a TestPrinter with the given answers as input.
func NewTestPrinterWithAnswers(answers []string) TestPrinter {
	return NewTestPrinterWithInput(bytes.NewBufferString(strings.Join(answers, "\n")))
}

// TestPrinter is an InputOutput implementation that uses a buffer for input.
type TestPrinter struct {
	stdPrinter
}

func (p TestPrinter) Outputs() TestOutputs {
	return p.InputOutput.(testInOut).TestOutputs //nolint:forcetypeassert
}

// TestOutputs is a struct that contains the output and error buffers.
type TestOutputs struct {
	Out, Err *bytes.Buffer
}

// OutOrStdout returns the output stream.
func (t TestOutputs) OutOrStdout() io.Writer {
	return t.Out
}

// ErrOrStderr returns the error stream.
func (t TestOutputs) ErrOrStderr() io.Writer {
	return t.Err
}

type testInOut struct {
	in io.Reader
	TestOutputs
}

func (t testInOut) InOrStdin() io.Reader {
	return t.in
}

var _ InputOutput = testInOut{}
