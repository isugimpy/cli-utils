package status

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var thirdPartyTypes = map[string]GetConditionsFn{
	"argoproj.io/Rollout": rolloutConditions,
}

// GetThirdPartyConditionsFn returns a function that can compute the status for the
// given resource, or nil if the resource type is not known.
func GetThirdPartyConditionsFn(u *unstructured.Unstructured) GetConditionsFn {
	gvk := u.GroupVersionKind()
	g := gvk.Group
	k := gvk.Kind
	key := g + "/" + k
	if g == "" {
		key = k
	}
	return thirdPartyTypes[key]
}

func rolloutConditions(u *unstructured.Unstructured) (*Result, error) {
	obj := u.UnstructuredContent()

	objc, err := GetObjectWithConditions(obj)
	if err != nil {
		return nil, err
	}

	completed := false
	healthy := false
	paused := false
	progressing := true
	progressingMessage := ""
	progressingReason := ""

	for _, c := range objc.Status.Conditions {
		switch c.Type {
		case "Completed":
			if c.Status == "True" {
				completed = true
			} else {
				completed = false
			}
		case "Healthy":
			if c.Status == "True" {
				healthy = true
			} else {
				healthy = false
			}
		case "Progressing":
			if c.Status == "True" {
				progressing = true
				progressingMessage = c.Message
			} else if c.Status == "Unknown" {
				progressing = false
			} else {
				if c.Reason == "ProgressDeadlineExceeded" {
					progressing = false
				}
			}
		case "Paused":
			if c.Status == "True" {
				paused = true
			} else {
				paused = false
			}
		}
	}

	if progressing {
		if !completed {
			return newInProgressStatus(progressingReason, progressingMessage), nil
		} else {
			if healthy {
				return &Result{
					Status:     CurrentStatus,
					Message:    "Rollout completed",
					Conditions: []Condition{},
				}, nil
			} else {
				return newFailedStatus(progressingReason, progressingMessage), nil
			}
		}
	} else {
		if completed && healthy {
			return &Result{
				Status:     CurrentStatus,
				Message:    "Rollout completed",
				Conditions: []Condition{},
			}, nil
		} else {
			if paused {
				return newInProgressStatus(progressingReason, progressingMessage), nil
			}
			return newFailedStatus(progressingReason, progressingMessage), nil
		}
	}
}
