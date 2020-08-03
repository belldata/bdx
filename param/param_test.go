package param_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/belldata/bdx/param"
)

func TestParams(t *testing.T) {
	pm := param.Params{
		param.Param{
			Key:   "id",
			Value: "1",
		},
	}
	id := pm.ByName("id")
	assert.Equal(t, id, "1")
}
