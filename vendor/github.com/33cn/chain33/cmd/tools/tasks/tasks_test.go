package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/33cn/chain33/util"
)

func TestReplaceTarget(t *testing.T) {
	fileName := "../config/template/executor/${CLASSNAME}.go.tmp"
	bcontent, err := util.ReadFile(fileName)
	assert.NoError(t, err)
	t.Log(string(bcontent))
}
