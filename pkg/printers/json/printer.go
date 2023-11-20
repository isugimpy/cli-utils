// Copyright 2020 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package json

import (
	"github.com/fluxcd/cli-utils/pkg/common"
	"github.com/fluxcd/cli-utils/pkg/print/list"
	"github.com/fluxcd/cli-utils/pkg/printers/printer"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func NewPrinter(ioStreams genericclioptions.IOStreams) printer.Printer {
	return &list.BaseListPrinter{
		FormatterFactory: func(previewStrategy common.DryRunStrategy) list.Formatter {
			return NewFormatter(ioStreams, previewStrategy)
		},
	}
}
