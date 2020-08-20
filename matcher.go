package matcher

import (
	"context"
	"reflect"
	"time"
)

// TimeWithGracePeriod will return a testify mock argument matcher that has a time.Duration argument
// for how much of a grace period should be given. For example if the given duration is 1*time.Minute it will
// allow time.Time value that are 1 minute before and 1 minute after the given time.
func TimeWithGracePeriod(duration time.Duration) func(time.Time) bool {
	return func(t time.Time) bool {
		return t.After(t.Add(-duration)) && t.Before(t.Add(duration))
	}
}

// ContextWithValue will return a testify mock argument matcher that has will match context.Context
// and compare a given value stored in the context with the given argument ctxValue
func ContextWithValue(key string, ctxValue interface{}) func(ctx context.Context) bool {
	return func(ctx context.Context) bool {
		return reflect.DeepEqual(ctx.Value(key), ctxValue)
	}
}
