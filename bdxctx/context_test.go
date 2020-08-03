package bdxctx_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/belldata/bdx/interfaces"
)

func TestHandlersChainEqual(t *testing.T) {
	var handler interfaces.BdxHandlerFunc
	handler = func(interfaces.Context) {}
	handlers := interfaces.HandlersChain{
		func(interfaces.Context) {},
		handler,
	}
	last := handlers.Last()
	val1 := reflect.ValueOf(handler)
	val2 := reflect.ValueOf(last)
	assert.Equal(t, val1, val2)
}
func TestHandlersChainNotEqual(t *testing.T) {
	var handler interfaces.BdxHandlerFunc
	handler = func(interfaces.Context) {}
	handlers := interfaces.HandlersChain{
		handler,
		func(interfaces.Context) {},
	}
	last := handlers.Last()
	val1 := reflect.ValueOf(handler)
	val2 := reflect.ValueOf(last)
	assert.NotEqual(t, val1, val2)
}

func TestContext(t *testing.T) {
	// ctx := bdxctx.New(nil, nil)
}
