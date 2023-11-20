// Copyright 2020 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package printers

import (
	"github.com/fluxcd/cli-utils/cmd/status/printers/event"
	"github.com/fluxcd/cli-utils/cmd/status/printers/json"
	"github.com/fluxcd/cli-utils/cmd/status/printers/printer"
	"github.com/fluxcd/cli-utils/cmd/status/printers/table"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// CreatePrinter return an implementation of the Printer interface. The
// actual implementation is based on the printerType requested.
func CreatePrinter(printerType string, ioStreams genericclioptions.IOStreams, printData *printer.PrintData) (printer.Printer, error) {
	switch printerType {
	case "table":
		return table.NewPrinter(ioStreams, printData), nil
	case "json":
		return json.NewPrinter(ioStreams, printData), nil
	default:
		return event.NewPrinter(ioStreams, printData), nil
	}
}
