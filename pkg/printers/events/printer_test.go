// Copyright 2020 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package events

import (
	"testing"

	"github.com/fluxcd/cli-utils/pkg/printers/printer"
	printertesting "github.com/fluxcd/cli-utils/pkg/printers/testutil"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func TestPrint(t *testing.T) {
	printertesting.PrintResultErrorTest(t, func() printer.Printer {
		ioStreams, _, _, _ := genericclioptions.NewTestIOStreams()
		return NewPrinter(ioStreams)
	})
}
