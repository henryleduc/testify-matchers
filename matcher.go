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
func ContextWithValue(key interface{}, ctxValue interface{}) func(ctx context.Context) bool {
	return func(ctx context.Context) bool {
		return reflect.DeepEqual(ctx.Value(key), ctxValue)
	}
}

// AnyContext will return a testify mock argument matcher that will match any context.Context implementation
// this allows the testing code to not care what happens to the context, while retaining runtime type guarantees in tests
func AnyContext() func(interface{}) bool {
	return func(i interface{}) bool {
		return AnythingWithInterface(compileTimeContextReference)(i)
	}
}

func compileTimeContextReference(_ context.Context) {}

// AnythingWithInterface expects a single parameter of type `func (_ myInterface){}`,
// and will return a matcher to match any given interface{} against your desired Interface
//
// I hate it.
// - There's enough logic going on here that it needs to be able to return an error,
//but it's trying to implement a matcher interface that doesn't support an error
// - there's a solid chance for panicking because of hacky reflections.  probably won't happen.
// - errors are masked and returned in-band because the default return type is the error scenario
//(i.e. we never explicitly return false, but we catch a panic and return before possibly returning true)
// - the function signature flies in the face of the type system and makes you read this big nasty comment
//
// why was it implemented this way?  Because I know of no way to pass an Interface.  I'm vaguely certain
// it's not possible due to how go handles its interfaces for compile/linking time guarantees
//(see runtime/runtime2.go's itab definition for more details)
func AnythingWithInterface(funcIfaceRef interface{}) func(toMatch interface{}) bool {
	return func(toMatch interface{}) bool {

		defer recover() // incredibly hacky, but so is the rest of this code

		targetType := reflect.TypeOf(funcIfaceRef)
		if targetType == nil {
			return false
		}

		if targetType.Kind() != reflect.Func {
			return false
		}

		if targetType.NumIn() != 1 {
			return false
		}

		t := targetType.In(0)

		if t == nil || t.Kind() != reflect.Interface {
			return false
		}

		queryType := reflect.TypeOf(toMatch)

		if queryType == nil {
			return false
		}

		return queryType.Implements(t)
	}
}
