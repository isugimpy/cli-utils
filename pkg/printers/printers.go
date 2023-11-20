// Copyright 2020 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package printers

import (
	"github.com/fluxcd/cli-utils/pkg/common"
	"github.com/fluxcd/cli-utils/pkg/print/list"
	"github.com/fluxcd/cli-utils/pkg/printers/events"
	"github.com/fluxcd/cli-utils/pkg/printers/json"
	"github.com/fluxcd/cli-utils/pkg/printers/printer"
	"github.com/fluxcd/cli-utils/pkg/printers/table"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

const (
	EventsPrinter = "events"
	TablePrinter  = "table"
	JSONPrinter   = "json"
)

func GetPrinter(printerType string, ioStreams genericclioptions.IOStreams) printer.Printer {
	switch printerType { //nolint:gocritic
	case TablePrinter:
		return &table.Printer{
			IOStreams: ioStreams,
		}
	case JSONPrinter:
		return &list.BaseListPrinter{
			FormatterFactory: func(previewStrategy common.DryRunStrategy) list.Formatter {
				return json.NewFormatter(ioStreams, previewStrategy)
			},
		}
	default:
		return events.NewPrinter(ioStreams)
	}
}

func SupportedPrinters() []string {
	return []string{EventsPrinter, TablePrinter, JSONPrinter}
}

func DefaultPrinter() string {
	return EventsPrinter
}

func ValidatePrinterType(printerType string) bool {
	for _, p := range SupportedPrinters() {
		if printerType == p {
			return true
		}
	}
	return false
}
