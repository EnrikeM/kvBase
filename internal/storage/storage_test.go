package storage_test

import (
	"testing"

	"github.com/EnrikeM/kvBase/internal/storage"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestUnitStoragePositive(t *testing.T) {
	srvc := storage.NewService(zaptest.NewLogger(t))

	msg, err := srvc.Set([]string{"key", "val"})
	require.NoError(t, err)
	require.Equal(t, "key key set for val val", msg)

	msg, err = srvc.Get([]string{"key"})
	require.NoError(t, err)
	require.Equal(t, "val", msg)

	msg, err = srvc.Del([]string{"key"})
	require.NoError(t, err)
	require.Equal(t, "key deleted", msg)
}
