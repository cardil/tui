package output_test

import (
	"context"
	"testing"

	"github.com/cardil/tui/output"
	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	ctx := context.TODO()
	p := output.PrinterFrom(ctx)

	assert.NotNil(t, p)
	assert.Equal(t, output.OsPrinter, p)

	tp := output.NewTestPrinter()
	ctx = output.WithContext(ctx, tp)

	p = output.PrinterFrom(ctx)
	assert.NotNil(t, p)
	assert.Equal(t, tp, p)
}
