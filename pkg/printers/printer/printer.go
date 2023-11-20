// Copyright 2020 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package printer

import (
	"github.com/fluxcd/cli-utils/pkg/apply/event"
	"github.com/fluxcd/cli-utils/pkg/common"
)

type Printer interface {
	Print(ch <-chan event.Event, previewStrategy common.DryRunStrategy, printStatus bool) error
}
