package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptional_UnmarshalJSON(t *testing.T) {
	var o Optional[int]

	err := json.Unmarshal([]byte("42"), &o)
	require.NoError(t, err)
	require.True(t, o.Present)
	require.Equal(t, 42, o.Value)

	err = json.Unmarshal([]byte(" null "), &o)
	require.NoError(t, err)
	require.False(t, o.Present)
	require.Equal(t, 0, o.Value)
}

func TestOptional_MarshalJSON(t *testing.T) {
	b, err := json.Marshal(Optional[int]{
		Present: true,
		Value:   42,
	})
	require.NoError(t, err)
	require.JSONEq(t, "42", string(b))

	b, err = json.Marshal(Optional[string]{
		Present: true,
		Value:   "test",
	})
	require.NoError(t, err)
	require.JSONEq(t, `"test"`, string(b))

	b, err = json.Marshal(Optional[*int]{
		Present: true,
	})
	require.NoError(t, err)
	require.JSONEq(t, "null", string(b))

	b, err = json.Marshal(Optional[int]{
		Present: false,
		Value:   42,
	})
	require.NoError(t, err)
	require.JSONEq(t, "null", string(b))
}
