package output

import "context"

type printerKey struct{}

// PrinterFrom returns the Printer from the context, or the default printer if
// none is set. The default printer is OsPrinter.
func PrinterFrom(ctx context.Context) Printer {
	p, ok := ctx.Value(printerKey{}).(Printer)
	if !ok {
		return defaultPrinter()
	}
	return p
}

// WithContext returns a new context with the given Printer set.
func WithContext(ctx context.Context, p Printer) context.Context {
	return context.WithValue(ctx, printerKey{}, p)
}

func defaultPrinter() Printer {
	return OsPrinter
}
