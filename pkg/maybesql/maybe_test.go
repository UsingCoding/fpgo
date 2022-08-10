package maybesql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/UsingCoding/fpgo/pkg/maybe"
)

func TestMaybe_Scan(t *testing.T) {
	m := maybe.NewJust("date")

	sqlM, err := FromMaybe(m)
	assert.NoError(t, err)

	_ = sqlM
}
