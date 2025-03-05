package compute_test

import (
	"testing"

	"github.com/EnrikeM/kvBase/internal/compute"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestUnitComputePostitive(t *testing.T) {
	srvc := compute.NewService(zaptest.NewLogger(t), *compute.NewParserSrvc())
	query, err := srvc.Parse("SET key val")
	require.NoError(t, err)
	require.Equal(t, compute.Query{
		Method: compute.SET,
		Args:   []string{"key", "val"},
	}, query)

	query, err = srvc.Parse("GET key")
	require.NoError(t, err)
	require.Equal(t, compute.Query{
		Method: compute.GET,
		Args:   []string{"key"},
	}, query)

	query, err = srvc.Parse("DEL key")
	require.NoError(t, err)
	require.Equal(t, compute.Query{
		Method: compute.DEL,
		Args:   []string{"key"},
	}, query)
}
