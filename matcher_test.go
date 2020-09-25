package matcher

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestAnythingWithInterface(t *testing.T) {
	assert.False(t, AnythingWithInterface(compileTimeContextReference)(5))
	assert.False(t, AnythingWithInterface(1)(5))
	assert.False(t, AnythingWithInterface(1)(context.Background()))
	assert.False(t, AnythingWithInterface(nil)(nil))
	assert.False(t, AnythingWithInterface(1)(nil))
	assert.False(t, AnythingWithInterface(nil)(1))

	assert.True(t, AnythingWithInterface(compileTimeContextReference)(context.Background()))
	assert.True(t, AnythingWithInterface(compileTimeContextReference)(context.TODO()))
	assert.True(t, AnythingWithInterface(compileTimeContextReference)(context.WithValue(context.Background(), "foo", "bar")))

	assert.True(t, AnyContext()(context.Background()))
	assert.True(t, AnyContext()(context.TODO()))
}

func TestAnyContext(t *testing.T) {
	assert.True(t, AnyContext()(context.Background()))
	assert.True(t, AnyContext()(context.TODO()))
	c, _ := context.WithCancel(context.Background())
	assert.True(t, AnyContext()(c))
	assert.True(t, AnyContext()(context.WithValue(c, "foo", "bar")))

	assert.False(t, AnyContext()(nil))
	assert.False(t, AnyContext()(""))
	assert.False(t, AnyContext()(0))
}