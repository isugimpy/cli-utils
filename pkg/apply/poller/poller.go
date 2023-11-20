// Copyright 2020 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package poller

import (
	"context"

	"github.com/fluxcd/cli-utils/pkg/kstatus/polling"
	pollevent "github.com/fluxcd/cli-utils/pkg/kstatus/polling/event"
	"github.com/fluxcd/cli-utils/pkg/object"
)

// Poller defines the interface the applier needs to poll for status of resources.
// The context is the preferred way to shut down the poller.
// The identifiers defines the resources which the poller should poll and
// compute status for.
// The options allows callers to override some of the settings of the poller,
// like the polling frequency and the caching strategy.
type Poller interface {
	Poll(ctx context.Context, identifiers object.ObjMetadataSet, options polling.PollOptions) <-chan pollevent.Event
}
