package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_asString(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := asString("abc")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "abc", have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := asString(42)

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})
}
